package dto

type GameCompletedResponse struct {
	Message      string `json:"message"`
	WinnerTeamId string `json:"winnerTeamId"`
}
