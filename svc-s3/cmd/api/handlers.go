package main

import (
	"encoding/json"
	"net/http"
)

func (app *Config) svcS3(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Message: "you hit the s3 service",
	}

	out, _ := json.MarshalIndent(payload, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}


func (app *Config) svcS32(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Message: "you hit the s32 service",
	}

	out, _ := json.MarshalIndent(payload, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
