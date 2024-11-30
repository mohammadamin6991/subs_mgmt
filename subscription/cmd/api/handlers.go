package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gitlab.amin.run/general/project/subs-mgmt/subscription/internal/invoice"
	"gitlab.amin.run/general/project/subs-mgmt/subscription/internal/subscription"
)

func (app *Config) CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var inv invoice.Invoice
	if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdInv, err := app.InvoiceService.CreateInvoice(context.Background(), &inv)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating invoice: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdInv)
}

func (app *Config) GetInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	invoiceID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	inv, err := app.InvoiceService.GetInvoice(context.Background(), invoiceID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching invoice: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inv)
}

func (app *Config) UpdateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	invoiceID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	var inv invoice.Invoice
	if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	inv.ID = invoiceID
	updatedInv, err := app.InvoiceService.UpdateInvoice(context.Background(), &inv)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating invoice: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedInv)
}

func (app *Config) DeleteInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	invoiceID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	err = app.InvoiceService.DeleteInvoice(context.Background(), invoiceID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting invoice: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *Config) GetAllInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	invoices, err := app.InvoiceService.GetAllInvoices(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching invoices: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func (app *Config) CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var sub subscription.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdSub, err := app.SubscriptionService.CreateSubscription(context.Background(), &sub)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating subscription: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdSub)
}

func (app *Config) GetSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	subscriptionID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid subscription ID", http.StatusBadRequest)
		return
	}

	sub, err := app.SubscriptionService.GetSubscription(context.Background(), subscriptionID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching subscription: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)
}

func (app *Config) UpdateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	subscriptionID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid subscription ID", http.StatusBadRequest)
		return
	}

	var sub subscription.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	sub.ID = subscriptionID
	updatedSub, err := app.SubscriptionService.UpdateSubscription(context.Background(), &sub)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating subscription: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedSub)
}

func (app *Config) DeleteSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	subscriptionID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid subscription ID", http.StatusBadRequest)
		return
	}

	err = app.SubscriptionService.DeleteSubscription(context.Background(), subscriptionID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting subscription: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *Config) GetAllSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	subscriptions, err := app.SubscriptionService.GetAllSubscriptions(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching subscriptions: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptions)
}
