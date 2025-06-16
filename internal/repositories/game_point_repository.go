package repositories

import (
	"githup/Therocking/dominoes/internal/entities"

	"gorm.io/gorm"
)

type GamePointRepository interface {
	Create(gamePoint *entities.GamePoint) error
	FindByGameID(gameID string) ([]*entities.GamePoint, error)
	FindAllByGameAndTeamId(gameId, teamId string) ([]*entities.GamePoint, error)
}

type gamePointRepo struct {
	db *gorm.DB
}

func NewGamePointRepository(db *gorm.DB) GamePointRepository {
	return &gamePointRepo{db: db}
}

func (r *gamePointRepo) Create(gamePoint *entities.GamePoint) error {
	return r.db.Create(gamePoint).Error
}

func (r *gamePointRepo) FindByGameID(gameID string) ([]*entities.GamePoint, error) {
	var points []*entities.GamePoint
	err := r.db.Where("game_id = ?", gameID).Find(&points).Error
	return points, err
}

func (r *gamePointRepo) FindAllByGameAndTeamId(gameId, teamId string) ([]*entities.GamePoint, error) {
	var points []*entities.GamePoint
	err := r.db.Where("game_id = ? AND team_id = ?", gameId, teamId).Find(&points).Error
	return points, err
}
