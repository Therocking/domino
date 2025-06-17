package dto

type GamePointResponse struct {
	Id     string `json:"id"`
	Point  int    `json:"point"`
	GameId string `json:"gameId"`
	TeamId string `json:"teamId"`
}
