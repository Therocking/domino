package repositories

import (
	"githup/Therocking/dominoes/internal/entities"

	"gorm.io/gorm"
)

type ISessionRepository interface {
	Create(session *entities.Session) error
	FindByDeviceID(id string) (*entities.Session, error)
}

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) ISessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Create(session *entities.Session) error {
	return r.db.Create(session).Error
}

func (r *SessionRepository) FindByDeviceID(deviceId string) (*entities.Session, error) {
	var session entities.Session
	err := r.db.Preload("Teams").First(&session, "device_id = ?", deviceId).Error
	return &session, err
}
