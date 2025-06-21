// session_handler.go
package handlers

import (
	"encoding/json"
	"githup/Therocking/dominoes/internal/services"
	"githup/Therocking/dominoes/pkg"
	"net/http"
)

type SessionHandler struct {
	service services.ISessionService
}

func NewSessionHandler(service services.ISessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

func (h *SessionHandler) CreateSession(response http.ResponseWriter, request *http.Request) {
	var dto struct {
		DeviceID string `json:"deviceId"`
	}

	if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
		pkg.WriteError(response, http.StatusBadRequest, err.Error())
		return
	}

	session, err := h.service.CreateSession(dto.DeviceID)
	if err != nil {
		pkg.WriteError(response, http.StatusBadRequest, err.Error())
		return
	}

	pkg.WriteJSON(response, http.StatusCreated, session)
}

func (h *SessionHandler) GetByDeviceId(response http.ResponseWriter, request *http.Request) {
	var deviceId = request.PathValue("deviceId")

	session, err := h.service.GetByDeviceId(deviceId)

	if err != nil {
		pkg.WriteError(response, http.StatusBadRequest, err.Error())
		return
	}

	pkg.WriteJSON(response, http.StatusOK, session)
}
