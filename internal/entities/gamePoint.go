package entities

type GamePoint struct {
	Base
	Point  int    `gorm:"not null" json:"point"`
	GameID string `gorm:"type:uuid;not null" json:"game_id"`
	TeamID string `gorm:"type:uuid;not null" json:"team_id"`
	Game   *Game  `gorm:"foreignKey:GameID"`
	Team   *Team  `gorm:"foreignKey:TeamID"`
}
