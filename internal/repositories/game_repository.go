package repositories

import (
	"githup/Therocking/dominoes/internal/entities"
	"time"

	"gorm.io/gorm"
)

type GameRepository interface {
	Create(game *entities.Game) error
	Update(game *entities.Game) error
	FindByID(id string) (*entities.Game, error)
}

type gameRepo struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return &gameRepo{db: db}
}

func (r *gameRepo) Create(game *entities.Game) error {
	return r.db.Create(game).Error
}

func (r *gameRepo) Update(game *entities.Game) error {
	now := time.Now()

	game.UpdatedAt = &now
	return r.db.Save(game).Error
}

func (r *gameRepo) FindByID(id string) (*entities.Game, error) {
	var game entities.Game
	err := r.db.First(&game, "id = ?", id).Error
	return &game, err
}
