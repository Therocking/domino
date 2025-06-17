package repositories

import (
	"githup/Therocking/dominoes/internal/entities"

	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(session *entities.Session) error
	FindByDeviceID(id string) (*entities.Session, error)
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepo{db: db}
}

func (r *sessionRepo) Create(session *entities.Session) error {
	return r.db.Create(session).Error
}

func (r *sessionRepo) FindByDeviceID(deviceId string) (*entities.Session, error) {
	var session entities.Session
	err := r.db.Preload("Teams").First(&session, "device_id = ?", deviceId).Error
	return &session, err
}
