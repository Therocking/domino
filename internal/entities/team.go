package entities

type Team struct {
	Base
	Name       string       `gorm:"not null" json:"name"`
	SessionID  string       `gorm:"type:uuid;not null" json:"session_id"`
	GameID     string       `gorm:"type:uuid;not null" json:"game_id"`
	Game       *Game        `gorm:"foreignKey:GameID"`
	GamePoints []*GamePoint `gorm:"foreignKey:TeamID"`
	Rankings   []*Ranking   `gorm:"foreignKey:TeamID"`
}
