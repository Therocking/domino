package repositories

import (
	"githup/Therocking/dominoes/internal/entities"

	"gorm.io/gorm"
)

type RankingRepository interface {
	Upsert(ranking *entities.Ranking) error
	FindAllByTeam(teamID string) ([]*entities.Ranking, error)
	FindByTeam(teamID string) (*entities.Ranking, error)
}

type rankingRepo struct {
	db *gorm.DB
}

func NewRankingRepository(db *gorm.DB) RankingRepository {
	return &rankingRepo{db: db}
}

func (r *rankingRepo) Upsert(ranking *entities.Ranking) error {
	return r.db.Save(ranking).Error
}

func (r *rankingRepo) FindAllByTeam(teamID string) ([]*entities.Ranking, error) {
	var rankings []*entities.Ranking
	err := r.db.Where("team_id = ?", teamID).Find(&rankings).Error
	return rankings, err
}

func (r *rankingRepo) FindByTeam(teamId string) (*entities.Ranking, error) {
	var ranking entities.Ranking
	err := r.db.Find(&ranking, "team_id = ?", teamId).Error
	return &ranking, err
}
