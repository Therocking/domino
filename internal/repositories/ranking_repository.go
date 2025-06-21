package repositories

import (
	"githup/Therocking/dominoes/internal/entities"

	"gorm.io/gorm"
)

type IRankingRepository interface {
	Upsert(ranking *entities.Ranking) error
	FindAllByTeam(teamID string) ([]*entities.Ranking, error)
	FindByTeam(teamID string) (*entities.Ranking, error)
}

type RankingRepository struct {
	db *gorm.DB
}

func NewRankingRepository(db *gorm.DB) IRankingRepository {
	return &RankingRepository{db: db}
}

func (r *RankingRepository) Upsert(ranking *entities.Ranking) error {
	return r.db.Save(ranking).Error
}

func (r *RankingRepository) FindAllByTeam(teamID string) ([]*entities.Ranking, error) {
	var rankings []*entities.Ranking
	err := r.db.Where("team_id = ?", teamID).Find(&rankings).Error
	return rankings, err
}

func (r *RankingRepository) FindByTeam(teamId string) (*entities.Ranking, error) {
	var ranking entities.Ranking
	err := r.db.Find(&ranking, "team_id = ?", teamId).Error
	return &ranking, err
}
