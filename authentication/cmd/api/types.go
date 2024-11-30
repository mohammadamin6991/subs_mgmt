package main

import "time"

type UserReq struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}

type UserRes struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin"`
}

type ListUserRes struct {
	Users []UserRes `json:"users"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User                  UserRes   `json:"user"`
}

type RenewAccessTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewAccessTokenRes struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type ValidationTokenRes struct {
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}
