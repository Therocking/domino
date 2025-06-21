package repositories

import (
	"githup/Therocking/dominoes/internal/entities"
	"time"

	"gorm.io/gorm"
)

type ITeamRepository interface {
	Create(team *entities.Team) error
	Update(team *entities.Team) error
	FindByID(id string) (*entities.Team, error)
	FindBySessionID(sessionID string) ([]*entities.Team, error)
	FindByGameID(gameID string) ([]*entities.Team, error)
}

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) ITeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) Create(team *entities.Team) error {
	return r.db.Create(team).Error
}

func (r *TeamRepository) Update(team *entities.Team) error {
	now := time.Now()

	team.UpdatedAt = &now
	return r.db.Save(team).Error
}

func (r *TeamRepository) FindByID(id string) (*entities.Team, error) {
	var team entities.Team
	err := r.db.First(&team, "id = ?", id).Error
	return &team, err
}

func (r *TeamRepository) FindBySessionID(sessionID string) ([]*entities.Team, error) {
	var teams []*entities.Team
	err := r.db.Where("session_id = ?", sessionID).Find(&teams).Error
	return teams, err
}

func (r *TeamRepository) FindByGameID(gameID string) ([]*entities.Team, error) {
	var teams []*entities.Team
	err := r.db.Where("game_id = ?", gameID).Find(&teams).Error
	return teams, err
}
