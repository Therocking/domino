package api

import (
	"githup/Therocking/dominoes/api/handlers"
	"githup/Therocking/dominoes/internal/repositories"
	"githup/Therocking/dominoes/internal/services"
	"githup/Therocking/dominoes/pkg/database"
	"net/http"
)

type RouterRegister struct{}

func (routerRegister *RouterRegister) GameRoutes(router *http.ServeMux) {
	gameRepo := repositories.NewGameRepository(database.DB)
	gamePointRepo := repositories.NewGamePointRepository(database.DB)
	rankingRepo := repositories.NewRankingRepository(database.DB)

	gameService := services.NewGameService(gameRepo, gamePointRepo, rankingRepo)

	gameHandler := handlers.NewGameHandler(gameService)

	router.HandleFunc("GET /game/{id}", gameHandler.GetPointsByGameId)
	router.HandleFunc("POST /game", gameHandler.AddPoint)
}

func (routerRegister *RouterRegister) SessionRoutes(router *http.ServeMux) {
	sessionRepo := repositories.NewSessionRepository(database.DB)
	teamRepo := repositories.NewTeamRepository(database.DB)
	gameRepo := repositories.NewGameRepository(database.DB)
	rankingRepo := repositories.NewRankingRepository(database.DB)

	sessionService := services.NewSessionService(sessionRepo, gameRepo, teamRepo)
	teamService := services.NewTeamService(teamRepo, rankingRepo)

	sessionHandler := handlers.NewSessionHandler(sessionService)
	teamHandler := handlers.NewTeamHandler(teamService)

	router.HandleFunc("POST /session", sessionHandler.CreateSession)
	router.HandleFunc("GET /session/{id}/teams", teamHandler.GetTeamsBySession)
	router.HandleFunc("GET /session/device/id/{deviceId}", sessionHandler.GetByDeviceId)
}

func (routerRegister *RouterRegister) TeamRoutes(router *http.ServeMux) {
	teamRepo := repositories.NewTeamRepository(database.DB)
	rankingRepo := repositories.NewRankingRepository(database.DB)

	teamService := services.NewTeamService(teamRepo, rankingRepo)

	teamHandler := handlers.NewTeamHandler(teamService)

	router.HandleFunc("GET /team/{id}/ranking", teamHandler.GetRanking)
	router.HandleFunc("GET /team/game/{gameId}/ranking", teamHandler.GetTeamsByGame)
	router.HandleFunc("PATCH /team/{id}", teamHandler.UpdateTeamName)
}
