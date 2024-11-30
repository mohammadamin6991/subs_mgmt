package main

import (
	"fmt"
	"log"
	"net/http"
)

const webport = "80"
type Config struct {

}

func main() {
	app := Config{}

	log.Printf("Starting svc-s3 service on port %s\n", webport)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webport),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
