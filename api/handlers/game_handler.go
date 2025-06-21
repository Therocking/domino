package handlers

import (
	"encoding/json"
	dto "githup/Therocking/dominoes/internal/dtos/game"
	"githup/Therocking/dominoes/internal/services"
	"githup/Therocking/dominoes/pkg"
	"net/http"
)

type GameHandler struct {
	service services.IGameService
}

func NewGameHandler(service services.IGameService) *GameHandler {
	return &GameHandler{service: service}
}

func (h *GameHandler) AddPoint(response http.ResponseWriter, request *http.Request) {
	var dto dto.CreateGamePoint

	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		pkg.WriteError(response, http.StatusInternalServerError, err.Error())
		return
	}

	pointResponse, err := h.service.AddPoint(&dto)

	if err != nil {
		pkg.WriteError(response, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteJSON(response, http.StatusOK, pointResponse)
}

func (h *GameHandler) GetPointsByGameId(response http.ResponseWriter, request *http.Request) {
	gameId := request.PathValue("id")

	gamePoint, err := h.service.GetPointsByGameId(gameId)

	if err != nil {
		pkg.WriteError(response, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.WriteJSON(response, http.StatusOK, gamePoint)
}
