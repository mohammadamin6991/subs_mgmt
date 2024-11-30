package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gitlab.amin.run/general/project/subs-mgmt/subscription/db"
	"gitlab.amin.run/general/project/subs-mgmt/subscription/internal/invoice"
	"gitlab.amin.run/general/project/subs-mgmt/subscription/internal/subscription"
)

const webport = "80"

type Config struct {
	InvoiceService      *invoice.InvoiceService
	SubscriptionService *subscription.SubscriptionService
}

var count int16

func main() {

	app := Config{}
	db := connectToDB()

	// Initialize repositories
	invoiceRepository := invoice.NewInvoiceRepository(db.GetDB())
	subscriptionRepository := subscription.NewSubscriptionRepository(db.GetDB())

	// Initialize services
	invoiceService := invoice.NewInvoiceService(invoiceRepository)
	subscriptionService := subscription.NewSubscriptionService(subscriptionRepository)

	app.InvoiceService = invoiceService
	app.SubscriptionService = subscriptionService

	log.Printf("Starting subscription service on port %s\n", webport)

	// Define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webport),
		Handler: app.routes(),
	}

	// Start the server
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func connectToDB() *db.Database {
	dsn := os.Getenv("DSN")
	for {
		connection, err := db.NewDatabase(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			count++
		} else {
			log.Println("Connected to Postgres Successfully")
			return connection
		}

		if count > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
	}
}
