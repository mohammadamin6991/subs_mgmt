package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gitlab.amin.run/general/project/subs-mgmt/authentication/internal/repository"
	"gitlab.amin.run/general/project/subs-mgmt/authentication/internal/service"
	"gitlab.amin.run/general/project/subs-mgmt/authentication/internal/token"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = 80

// ??? Why not define it in the function it self?
var count int16

type Config struct {
	UserService *service.UserService
	SessionService *service.SessionService
	Context context.Context
	TokenMaker *token.JWTMaker
}

func main() {
	log.Println("Starting authentication service...")

	conn := connectToDB()
	userRepository := repository.NewUserRepository(conn.GetDB())
	sessionRepository := repository.NewSessionRepository(conn.GetDB())
	userService := service.NewUserService(userRepository)
	sessionService := service.NewSessionService(sessionRepository)

	tokenMaker := token.NewJWTMaker(os.Getenv("JWT_TOKEN"))

	if conn == nil {
		log.Panic("Can not connect to Postgres")
	}
	// Set-up config
	app := Config{
		UserService: userService,
		SessionService: sessionService,
		Context: context.Background(),
		TokenMaker: tokenMaker,
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", strconv.Itoa(webPort)),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func connectToDB() *repository.Database {
	dsn := os.Getenv("DSN")
	for {
		connection, err := repository.NewDatabase(dsn)
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
