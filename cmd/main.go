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

	sessionGroup := router.Group("/api/v1/sessions")
	{
		sessionGroup.POST("/", sessionHandler.CreateSession)
		sessionGroup.GET("/device/:deviceId", sessionHandler.GetByDeviceId)
		sessionGroup.GET("/:sessionId/teams", teamHandler.GetTeamsBySession)
	}

	gameGroup := router.Group("/api/v1/games")
	{
		gameGroup.GET("/:gameId/teams", teamHandler.GetTeamsByGame)
		gameGroup.POST("/points", gameHandler.AddPoint)
		gameGroup.GET("/:gameId/points", gameHandler.GetPointsByGameId)
	}

	teamGroup := router.Group("/api/v1/team")
	{

		teamGroup.GET("/:teamId/ranking", teamHandler.GetRanking)
		teamGroup.PATCH("/:teamId", teamHandler.UpdateTeamName)
	}

	// Iniciar servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
