package main

import (
	"encoding/json"
	"errors"
	"net/http"

	utils "gitlab.amin.run/general/project/subs-mgmt/gateway/pkg"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Gateway(w http.ResponseWriter, r *http.Request) {
	payload := utils.JsonResponse{
		Error:   false,
		Message: "Hit the gateway",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	w.Write(out)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err)
	}

	switch requestPayload.Action {
	case "auth":
	default:
		utils.ErrorJSON(w, errors.New("unknown action"))
	}
}


func (app *Config) adminRoleReq(w http.ResponseWriter, r *http.Request) {
	payload := utils.JsonResponse{
		Error:   false,
		Message: "Admin Role Req",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	w.Write(out)
}

func (app *Config) normalRoleReq(w http.ResponseWriter, r *http.Request) {
	payload := utils.JsonResponse{
		Error:   false,
		Message: "Normal User",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	w.Write(out)
}

func (app *Config) openReq(w http.ResponseWriter, r *http.Request) {
	payload := utils.JsonResponse{
		Error:   false,
		Message: "Open",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	w.Write(out)
}
