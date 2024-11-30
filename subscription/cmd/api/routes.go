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
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/healthz"))

	// Plan Routes
	mux.Post("/subscriptions", app.CreateSubscriptionHandler)
	mux.Get("/subscriptions/{id}", app.GetSubscriptionHandler)
	mux.Put("/subscriptions/{id}", app.UpdateSubscriptionHandler)
	mux.Delete("/subscriptions/{id}", app.DeleteSubscriptionHandler)
	mux.Get("/subscriptions", app.GetAllSubscriptionsHandler)

	mux.Post("/invoices", app.CreateInvoiceHandler)
	mux.Get("/invoices/{id}", app.GetInvoiceHandler)
	mux.Put("/invoices/{id}", app.UpdateInvoiceHandler)
	mux.Delete("/invoices/{id}", app.DeleteInvoiceHandler)
	mux.Get("/invoices", app.GetAllInvoicesHandler)

	return mux
}
