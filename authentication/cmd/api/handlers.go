package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gitlab.amin.run/general/project/subs-mgmt/authentication/internal/repository"
	utils "gitlab.amin.run/general/project/subs-mgmt/authentication/pkg"

	"github.com/go-chi/chi/v5"
)

func (app *Config) createUser(w http.ResponseWriter, r *http.Request) {
	var u UserReq
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// hash password
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}
	u.Password = hashed

	created, err := app.UserService.Create(app.Context, toStorerUser(u))
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, fmt.Sprintf("error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	res := toUserRes(created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func toStorerUser(u UserReq) *repository.User {
	return &repository.User{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Password:  u.Password,
		IsAdmin:   u.IsAdmin,
	}
}

func toUserRes(u *repository.User) UserRes {
	return UserRes{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		IsAdmin:   u.IsAdmin,
	}
}

func (app *Config) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.UserService.List(app.Context)
	if err != nil {
		http.Error(w, fmt.Sprintf("error listing users: %v", err), http.StatusInternalServerError)
		return
	}

	var res ListUserRes
	for _, u := range users {
		res.Users = append(res.Users, toUserRes(&u))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (app *Config) updateUser(w http.ResponseWriter, r *http.Request) {
	// we will later get user email from the token payload of the authenticated user
	var u UserReq
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	user, err := app.UserService.GetByEmail(app.Context, u.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	// patch our user request
	patchUserReq(user, u)

	updated, err := app.UserService.Update(app.Context, user)
	if err != nil {
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}

	res := toUserRes(updated)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func patchUserReq(user *repository.User, u UserReq) {
	if u.Firstname != "" {
		user.Firstname = u.Firstname
	}
	if u.Lastname != "" {
		user.Lastname = u.Lastname
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.Password != "" {
		hashed, err := utils.HashPassword(u.Password)
		if err != nil {
			panic(err)
		}
		user.Password = hashed
	}
	if u.IsAdmin {
		user.IsAdmin = u.IsAdmin
	}
	user.UpdatedAt = toTimePtr(time.Now())
}

func toTimePtr(t time.Time) *time.Time {
	return &t
}

func (app *Config) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", http.StatusBadRequest)
		return
	}

	err = app.UserService.Delete(app.Context, i)
	if err != nil {
		http.Error(w, "error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *Config) loginUser(w http.ResponseWriter, r *http.Request) {
	var u LoginUserReq
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	gu, err := app.UserService.GetByEmail(app.Context, u.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting user: %v", err), http.StatusInternalServerError)
		return
	}

	err = utils.CheckPassword(u.Password, gu.Password)
	if err != nil {
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	}

	// create a json web token (JWT) and return it as response
	accessToken, accessClaims, err := app.TokenMaker.CreateToken(gu.ID, gu.Email, gu.IsAdmin, 15*time.Minute)
	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	refreshToken, refreshClaims, err := app.TokenMaker.CreateToken(gu.ID, gu.Email, gu.IsAdmin, 24*time.Hour)
	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	session, err := app.SessionService.CreateSession(app.Context, &repository.Session{
		ID:           refreshClaims.RegisteredClaims.ID,
		UserEmail:    gu.Email,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    refreshClaims.RegisteredClaims.ExpiresAt.Time,
	})
	if err != nil {
		http.Error(w, "error creating session", http.StatusInternalServerError)
		return
	}

	res := LoginUserRes{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.RegisteredClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time,
		User:                  toUserRes(gu),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (app *Config) logoutUser(w http.ResponseWriter, r *http.Request) {
	// we will later get the session ID from the token payload of the authenticated user
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing session ID", http.StatusBadRequest)
		return
	}

	err := app.SessionService.DeleteSession(app.Context, id)
	if err != nil {
		http.Error(w, "error deleting session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *Config) validateToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "bad auth header", http.StatusBadRequest)
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 || fields[0] != "Bearer" {
		http.Error(w, "invalid authorization header", http.StatusBadRequest)
	}

	token := fields[1]
	claims, err := app.TokenMaker.VerifyToken(token)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid token: %v", err), http.StatusBadRequest)
	}

	res := ValidationTokenRes{
		Email:             claims.Email,
		IsAdmin:           claims.IsAdmin,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (app *Config) renewAccessToken(w http.ResponseWriter, r *http.Request) {
	var req RenewAccessTokenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	refreshClaims, err := app.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "error verifying token", http.StatusUnauthorized)
		return
	}

	session, err := app.SessionService.GetSession(app.Context, refreshClaims.RegisteredClaims.ID)
	if err != nil {
		http.Error(w, "error getting session", http.StatusInternalServerError)
		return
	}

	if session.IsRevoked {
		http.Error(w, "session revoked", http.StatusUnauthorized)
		return
	}

	if session.UserEmail != refreshClaims.Email {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}

	accessToken, accessClaims, err := app.TokenMaker.CreateToken(refreshClaims.ID, refreshClaims.Email, refreshClaims.IsAdmin, 15*time.Minute)
	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	res := RenewAccessTokenRes{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (app *Config) revokeSession(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing session ID", http.StatusBadRequest)
		return
	}

	err := app.SessionService.RevokeSession(app.Context, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error revoking session: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
