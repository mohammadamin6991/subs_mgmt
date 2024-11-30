package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	middlewares "gitlab.amin.run/general/project/subs-mgmt/gateway/internal/middleware"
	"gitlab.amin.run/general/project/subs-mgmt/gateway/internal/proxy"
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

	// mux.Post("/", app.Gateway)

	// mux.Post("/", app.HandleSubmission)

	mux.Route("/nice", func(mux chi.Router) {
		mux.Post("/open", app.openReq)
	})

	mux.With(middlewares.WithAuth).Route("/", func(mux chi.Router) {
		mux.Post("/authenticated", app.normalRoleReq)

		mux.With(middlewares.WithAdminRole).Route("/admin", func(mux chi.Router) {
			mux.Post("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("POST request to /admin")
				proxy.SvcS3Proxy().ServeHTTP(w, r)
			})
			mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("GET request to /admin")
				proxy.SvcS3Proxy().ServeHTTP(w, r)
			})
		})
	})
	// Define routes for the subscription and invoice services
	mux.Post("/subscriptions", proxy.SubscriptionProxy().ServeHTTP)
	mux.Get("/subscriptions/{id}", proxy.SubscriptionProxy().ServeHTTP)
	mux.Put("/subscriptions/{id}", proxy.SubscriptionProxy().ServeHTTP)
	mux.Delete("/subscriptions/{id}",proxy.SubscriptionProxy().ServeHTTP)
	mux.Get("/subscriptions",proxy.SubscriptionProxy().ServeHTTP)

	mux.Post("/invoices",proxy.SubscriptionProxy().ServeHTTP)
	mux.Get("/invoices/{id}",proxy.SubscriptionProxy().ServeHTTP)
	mux.Put("/invoices/{id}",proxy.SubscriptionProxy().ServeHTTP)
	mux.Delete("/invoices/{id}",proxy.SubscriptionProxy().ServeHTTP)
	mux.Get("/invoices",proxy.SubscriptionProxy().ServeHTTP)
	// Proxy routes
	mux.Route("/admin", func(mux chi.Router) {
		// Authentication API Proxy
		mux.Post("/auth", proxy.SvcS3Proxy())
		mux.Get("/auth", proxy.SvcS3Proxy())

		// S3 Proxy
		mux.Post("/svc-s3", proxy.SvcS3Proxy())
		mux.Get("/svc-s3", proxy.SvcS3Proxy())

		// Subscription Service Proxy
		mux.Post("/subscriptions", proxy.SubscriptionProxy().ServeHTTP)
		mux.Get("/subscriptions/{id}", proxy.SubscriptionProxy().ServeHTTP)
	})
	return mux
}
