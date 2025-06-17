package dto

import dto "githup/Therocking/dominoes/internal/dtos/team"

type SessionResponse struct {
	Id    string              `json:"id"`
	Teams []*dto.TeamResponse `json:"teams"`
}
