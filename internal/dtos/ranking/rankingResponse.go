package dto

type RankingResponse struct {
	Id          string  `json:"id"`
	GameId      string  `json:"gameId"`
	TeamId      string  `json:"teamId"`
	TotalGames  int     `json:"totalGames"`
	TotalWins   int     `json:"totalWins"`
	TotalLosses int     `json:"totalLosses"`
	WinRate     float64 `json:"winRate"`
}
