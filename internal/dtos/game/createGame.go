package dto

type CreateGame struct {
	Point  int    `json:"point"`
	GameId string `json:"gameId"`
	TeamId string `json:"teamId"`
}
