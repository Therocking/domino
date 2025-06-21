package main

import (
	"fmt"
	"githup/Therocking/dominoes/api"
	"githup/Therocking/dominoes/pkg/database"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	router := http.NewServeMux()
	routerRegister := &api.RouterRegister{}

	routerRegister.GameRoutes(router)
	routerRegister.SessionRoutes(router)
	routerRegister.TeamRoutes(router)

	apiRouter := http.NewServeMux()
	apiRouter.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: apiRouter,
	}

	log.Printf("Server running on port: %s", port)

	server.ListenAndServe()
}
