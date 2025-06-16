package entities

type Game struct {
	Base
	Completed bool         `gorm:"default:false" json:"completed"`
	Teams     []*Team      `gorm:"foreignKey:GameID"`
	Points    []*GamePoint `gorm:"foreignKey:GameID"`
}
