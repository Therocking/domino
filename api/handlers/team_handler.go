package handlers

import (
	"encoding/json"
	teamDto "githup/Therocking/dominoes/internal/dtos/team"
	"githup/Therocking/dominoes/internal/services"
	"githup/Therocking/dominoes/pkg"
	"net/http"
)

type TeamHandler struct {
	service services.TeamService
}

func NewTeamHandler(service services.TeamService) *TeamHandler {
	return &TeamHandler{service: service}
}

func (h *TeamHandler) GetTeamsBySession(response http.ResponseWriter, request *http.Request) {
	sessionID := request.PathValue("id")

	teams, err := h.service.GetTeamsBySession(sessionID)
	if err != nil {
		pkg.WriteError(response, http.StatusBadRequest, err.Error())
		return
	}

	pkg.WriteJSON(response, http.StatusOK, teams)
}

func (h *TeamHandler) GetTeamsByGame(response http.ResponseWriter, request *http.Request) {
	gameID := request.PathValue("gameId")

	teams, err := h.service.GetTeamsByGame(gameID)
	if err != nil {
		pkg.WriteError(response, http.StatusBadRequest, err.Error())
		return
	}

	pkg.WriteJSON(response, http.StatusOK, teams)
}

func (h *TeamHandler) GetRanking(response http.ResponseWriter, request *http.Request) {
	rankingId := request.PathValue("id")

	rankings, err := h.service.GetRanking(rankingId)
	if err != nil {
		pkg.WriteError(response, http.StatusBadRequest, err.Error())
		return
	}

	pkg.WriteJSON(response, http.StatusOK, rankings)
}

func (h *TeamHandler) UpdateTeamName(resonse http.ResponseWriter, request *http.Request) {
	teamId := request.PathValue("id")
	var dto *teamDto.UpdateTeamName

	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		pkg.WriteError(resonse, http.StatusBadRequest, err.Error())
		return
	}

	dto.Id = teamId

	err := h.service.UpdateTeamName(dto)
	if err != nil {
		pkg.WriteError(resonse, http.StatusBadRequest, err.Error())
		return
	}

	pkg.WriteJSON(resonse, http.StatusAccepted, nil)
}
