// session_handler.go
package handlers

import (
	"githup/Therocking/dominoes/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	service services.SessionService
}

func NewSessionHandler(service services.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

func (h *SessionHandler) CreateSession(c *gin.Context) {
	var request struct {
		DeviceID string `json:"deviceId"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := h.service.CreateSession(request.DeviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, session)
}
