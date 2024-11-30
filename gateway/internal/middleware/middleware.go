package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type AuthResponse struct {
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		authURL := os.Getenv("AUTH_ENDPOINT")
		req, err := http.NewRequest("POST", authURL, nil)
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating request: %v", err), http.StatusInternalServerError)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error sending request: %v", err), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()
		var ar AuthResponse
		log.Printf("body: %v", res.Body)
		if err := json.NewDecoder(res.Body).Decode(&ar); err != nil {
			http.Error(w, fmt.Sprintf("error decoding response body: %v\n %v", err, r.Body), http.StatusInternalServerError)
			return
		}
		if res.StatusCode == http.StatusOK {
			r.Header.Set("X-Username", ar.Email)
			r.Header.Set("X-IsAdmin", strconv.FormatBool(ar.IsAdmin))
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func WithAdminRole(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        isAdmin := r.Header.Get("X-IsAdmin") // Extract X-IsAdmin header

        if isAdmin != "true" {
            http.Error(w, "Forbidden: Admin role required", http.StatusForbidden)
            return
        }

        // If the user is an admin, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}
