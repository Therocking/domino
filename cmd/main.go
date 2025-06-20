package main

import (
	"githup/Therocking/dominoes/api"
	"githup/Therocking/dominoes/pkg/database"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	routerRegister := &api.RouterRegister{}

	routerRegister.GameRoutes(router)
	routerRegister.SessionRoutes(router)
	routerRegister.TeamRoutes(router)

	apiRouter := http.NewServeMux()
	apiRouter.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	server := http.Server{
		Addr:    ":8080",
		Handler: apiRouter,
	}

	server.ListenAndServe()
}
