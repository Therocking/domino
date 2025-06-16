package services

import (
	"errors"
	teamDto "githup/Therocking/dominoes/internal/dtos/team"
	"githup/Therocking/dominoes/internal/entities"
	"githup/Therocking/dominoes/internal/repositories"
)

type TeamService interface {
	UpdateTeamName(team *teamDto.UpdateTeamName) error
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

func (s *teamService) UpdateTeamName(dto *teamDto.UpdateTeamName) error {
	if dto.Name == "" {
		return errors.New("team name cannot be empty")
	}

	team, err := s.repo.FindByID(dto.Id)

	if err != nil {
		return err
	}

	team.Name = dto.Name
	err = s.repo.Update(team)

	if err != nil {
		return err
	}

	return nil
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
