package handlers

import (
	"githup/Therocking/dominoes/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	service services.TeamService
}

func NewTeamHandler(service services.TeamService) *TeamHandler {
	return &TeamHandler{service: service}
}

func (h *TeamHandler) GetTeamsBySession(c *gin.Context) {
	sessionID := c.Param("sessionId")

	teams, err := h.service.GetTeamsBySession(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *TeamHandler) GetTeamsByGame(c *gin.Context) {
	gameID := c.Param("gameId")

	teams, err := h.service.GetTeamsByGame(gameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *TeamHandler) GetRanking(c *gin.Context) {
	rankingId := c.Param("teamId")

	rankings, err := h.service.GetRanking(rankingId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rankings)
}
