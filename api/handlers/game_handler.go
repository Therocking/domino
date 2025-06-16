package handlers

import (
	dto "githup/Therocking/dominoes/internal/dtos/game"
	"githup/Therocking/dominoes/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	service services.GameService
}

func NewGameHandler(service services.GameService) *GameHandler {
	return &GameHandler{service: service}
}

func (h *GameHandler) AddPoint(c *gin.Context) {
	var dto dto.CreateGame
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddPoint(&dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *GameHandler) GetPointsByGameId(c *gin.Context) {
	gameId := c.Param("gameId")

	gamePoint, err := h.service.GetPointsByGameId(gameId)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gamePoint)
}
