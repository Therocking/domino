package services

import (
	"errors"
	"githup/Therocking/dominoes/internal/entities"
	"githup/Therocking/dominoes/internal/repositories"
)

type TeamService interface {
	CreateTeam(team *entities.Team) error
	GetTeamsBySession(sessionID string) ([]*entities.Team, error)
	GetTeamsByGame(gameID string) ([]*entities.Team, error)
	GetRanking(id string) ([]*entities.Ranking, error)
}

type teamService struct {
	repo        repositories.TeamRepository
	rankingRepo repositories.RankingRepository
}

func NewTeamService(repo repositories.TeamRepository, rankingRepo repositories.RankingRepository) TeamService {
	return &teamService{repo: repo, rankingRepo: rankingRepo}
}

func (s *teamService) CreateTeam(team *entities.Team) error {
	if team.Name == "" {
		return errors.New("team name cannot be empty")
	}
	return s.repo.Create(team)
}

func (s *teamService) GetTeamsBySession(sessionID string) ([]*entities.Team, error) {
	return s.repo.FindBySessionID(sessionID)
}

func (s *teamService) GetTeamsByGame(gameID string) ([]*entities.Team, error) {
	return s.repo.FindByGameID(gameID)
}

func (s *teamService) GetRanking(id string) ([]*entities.Ranking, error) {
	return s.rankingRepo.FindAllByTeam(id)
}
