package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const webPort = "80"

type Config struct {}

func main() {
	app := Config{}

	os.Getenv("LOG_LEVEL")
	os.Getenv("AUTH_ENDPOINT")
	os.Getenv("SVC_S3_ENDPOINT")
	os.Getenv("SUBSCRIPTION_ENDPOINT")

	log.Printf("Starting gateway service on port %s\n", webPort)

	// Define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
