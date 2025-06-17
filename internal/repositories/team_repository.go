package repositories

import (
	"githup/Therocking/dominoes/internal/entities"
	"time"

	"gorm.io/gorm"
)

type TeamRepository interface {
	Create(team *entities.Team) error
	Update(team *entities.Team) error
	FindByID(id string) (*entities.Team, error)
	FindBySessionID(sessionID string) ([]*entities.Team, error)
	FindByGameID(gameID string) ([]*entities.Team, error)
}

type teamRepo struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) Create(team *entities.Team) error {
	return r.db.Create(team).Error
}

func (r *teamRepo) Update(team *entities.Team) error {
	now := time.Now()

	team.UpdatedAt = &now
	return r.db.Save(team).Error
}

func (r *teamRepo) FindByID(id string) (*entities.Team, error) {
	var team entities.Team
	err := r.db.Preload("GamePoints").Preload("Rankings").First(&team, "id = ?", id).Error
	return &team, err
}

func (r *teamRepo) FindBySessionID(sessionID string) ([]*entities.Team, error) {
	var teams []*entities.Team
	err := r.db.Preload("GamePoints").Preload("Rankings").Where("session_id = ?", sessionID).Find(&teams).Error
	return teams, err
}

func (r *teamRepo) FindByGameID(gameID string) ([]*entities.Team, error) {
	var teams []*entities.Team
	err := r.db.Where("game_id = ?", gameID).Find(&teams).Error
	return teams, err
}
