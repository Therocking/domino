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

	sessionRepo := repositories.NewSessionRepository(db)
	teamRepo := repositories.NewTeamRepository(db)
	gameRepo := repositories.NewGameRepository(db)
	gamePointRepo := repositories.NewGamePointRepository(db)
	rankingRepo := repositories.NewRankingRepository(db)

	sessionService := services.NewSessionService(sessionRepo, gameRepo, teamRepo)
	gameService := services.NewGameService(gameRepo, gamePointRepo, rankingRepo)
	teamService := services.NewTeamService(teamRepo, rankingRepo)

	sessionHandler := handlers.NewSessionHandler(sessionService)
	gameHandler := handlers.NewGameHandler(gameService)
	teamHandler := handlers.NewTeamHandler(teamService)

	router := gin.Default()

	router.Use(CORSMiddleware())

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

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
