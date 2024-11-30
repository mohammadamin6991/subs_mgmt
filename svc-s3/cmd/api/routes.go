package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// Specify who can connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	mux.Use(middleware.Heartbeat("/healthz"))
	mux.Post("/", app.svcS3)
	mux.Get("/", app.svcS32)
	mux.Get("/admin", app.svcS32)

    // Plan Routes
    mux.Post("/plans",app.CreatePlanHandler)
    mux.Get("/plans",app.GetAllPlansHandler)
    mux.Get("/plans/{id}",app.GetPlanHandler)
    mux.Put("/plans/{id}",app.UpdatePlanHandler)
    mux.Delete("/plans/{id}",app.DeletePlanHandler)

    // Instance Routes
    mux.Post("/instances", app.CreateInstanceHandler)
    mux.Get("/instances", app.GetAllInstancesHandler)
    mux.Get("/instances/{id}", app.GetInstanceHandler)
    mux.Put("/instances/{id}", app.UpdateInstanceHandler)
    mux.Delete("/instances/{id}", app.DeleteInstanceHandler)

	return mux
}
