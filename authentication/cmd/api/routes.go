package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors. Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Route("/users", func(mux chi.Router) {
		mux.Post("/", app.createUser)
		mux.Get("/", app.listUsers)
		mux.Patch("/", app.updateUser)

		mux.Route("/{id}", func(mux chi.Router) {
			mux.Delete("/", app.deleteUser)
		})

		mux.Route("/login", func(mux chi.Router) {
			mux.Post("/", app.loginUser)
		})

		mux.Route("/logout", func(mux chi.Router) {
			mux.Post("/", app.logoutUser)
		})
	})

	mux.Route("/tokens", func(mux chi.Router) {
		mux.Route("/renew", func(mux chi.Router) {
			mux.Post("/", app.renewAccessToken)
		})
		mux.Route("/validate", func(mux chi.Router) {
			mux.Post("/", app.validateToken)
		})
		mux.Route("/revoke/{id}", func(mux chi.Router) {
			mux.Post("/", app.revokeSession)
		})
	})
	return mux
}
