package entities

import (
	"time"
)

type Base struct {
	ID           string     `gorm:"primaryKey;type:uuid"`
	CreatedAt    time.Time  `json:"created_at"`
	IsDeleted    bool       `json:"isDeleted"`
	DeletedToken *string    `json:"deletedToken"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
