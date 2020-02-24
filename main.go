package main

import (
	"fmt"
	"github.com/FernandoCagale/c4-notify/api/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	session, err := SetupMongoDB()
	if err != nil {
		fmt.Println(err.Error())
		panic("Erro to start MongoDB")
	}

	defer session.Close()

	app, err := SetupApplication(session)
	if err != nil {
		panic("Erro to start application")
	}

	app.MakeEvents()

	router := app.MakeHandlers()

	router.Use(middleware.Header)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
