package entities

type Session struct {
	Base
	DeviceID string  `gorm:"type:uuid;not null" json:"device_id"`
	Teams    []*Team `gorm:"foreignKey:SessionID"`
}
