package entities

type Ranking struct {
	Base
	TeamID      string  `gorm:"type:uuid;not null" json:"team_id"`
	GameID      string  `gorm:"type:uuid;not null" json:"game_id"`
	TotalGames  int     `gorm:"not null;default:0" json:"total_games"`
	TotalWins   int     `gorm:"not null;default:0" json:"total_wins"`
	TotalLosses int     `gorm:"not null;default:0" json:"total_losses"`
	WinRate     float64 `gorm:"not null;default:0" json:"win_rate"`
	Team        *Team   `gorm:"foreignKey:TeamID"`
	Game        *Game   `gorm:"foreignKey:GameID"`
}
