package main

import (
	"githup/Therocking/dominoes/api/handlers"
	"githup/Therocking/dominoes/internal/repositories"
	"githup/Therocking/dominoes/internal/services"
	"githup/Therocking/dominoes/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Inicializar repositorios
	sessionRepo := repositories.NewSessionRepository(db)
	teamRepo := repositories.NewTeamRepository(db)
	gameRepo := repositories.NewGameRepository(db)
	gamePointRepo := repositories.NewGamePointRepository(db)
	rankingRepo := repositories.NewRankingRepository(db)

	// Inicializar servicios
	sessionService := services.NewSessionService(sessionRepo, gameRepo, teamRepo)
	gameService := services.NewGameService(gameRepo, gamePointRepo, rankingRepo)
	teamService := services.NewTeamService(teamRepo, rankingRepo)

	// Inicializar handlers
	sessionHandler := handlers.NewSessionHandler(sessionService)
	gameHandler := handlers.NewGameHandler(gameService)
	teamHandler := handlers.NewTeamHandler(teamService)

	// Configurar router
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/sessions", sessionHandler.CreateSession)
		api.GET("/sessions/:sessionId/teams", teamHandler.GetTeamsBySession)
		api.GET("/games/:gameId/teams", teamHandler.GetTeamsByGame)
		api.POST("/games/points", gameHandler.AddPoint)
		api.GET("/games/:gameId/points", gameHandler.GetPointsByGameId)
		api.GET("/team/:teamId/ranking", teamHandler.GetRanking)
		api.PATCH("/team/:teamId", teamHandler.UpdateTeamName)
	}

	// Iniciar servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
