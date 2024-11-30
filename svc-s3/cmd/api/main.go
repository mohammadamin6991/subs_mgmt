package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gitlab.amin.run/general/project/subs-mgmt/svc-s3/db"
	"gitlab.amin.run/general/project/subs-mgmt/svc-s3/internal/instance"
	"gitlab.amin.run/general/project/subs-mgmt/svc-s3/internal/plan"
)

const webport = "80"

var count int16

type Config struct {
	PlanService     *plan.PlanService
	InstanceService *instance.InstanceService
}

func main() {

	app := Config{}
	db := connectToDB()

	// Initialize repositories
	planRepo := plan.NewPlanRepository(db.GetDB())
	instanceRepo := instance.NewInstanceRepository(db.GetDB())

	// Initialize services
	planService := plan.NewPlanService(planRepo)
	instanceService := instance.NewInstanceService(instanceRepo)

	app.PlanService = planService
	app.InstanceService = instanceService

	log.Printf("Starting svc-s3 service on port %s\n", webport)

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
