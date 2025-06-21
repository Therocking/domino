package services

import (
	"errors"
	rankingDto "githup/Therocking/dominoes/internal/dtos/ranking"
	teamDto "githup/Therocking/dominoes/internal/dtos/team"
	"githup/Therocking/dominoes/internal/entities"
	"githup/Therocking/dominoes/internal/repositories"
)

type ITeamService interface {
	UpdateTeamName(team *teamDto.UpdateTeamName) error
	GetTeamsBySession(sessionID string) ([]*teamDto.TeamResponse, error)
	GetTeamsByGame(gameID string) ([]*teamDto.TeamResponse, error)
	GetRanking(id string) ([]*rankingDto.RankingResponse, error)
}

type TeamService struct {
	repo        repositories.ITeamRepository
	rankingRepo repositories.IRankingRepository
}

func NewTeamService(repo repositories.ITeamRepository, rankingRepo repositories.IRankingRepository) ITeamService {
	return &TeamService{repo: repo, rankingRepo: rankingRepo}
}

func (s *TeamService) UpdateTeamName(dto *teamDto.UpdateTeamName) error {
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

func (s *TeamService) GetTeamsBySession(sessionID string) ([]*teamDto.TeamResponse, error) {
	teams, err := s.repo.FindBySessionID(sessionID)

	var response []*teamDto.TeamResponse

	if err != nil {
		return response, err
	}

	for i := 0; i < len(teams); i++ {
		team := teams[i]

		teamTransformed := s.transformResponse(team)

		response = append(response, teamTransformed)
	}

	return response, nil
}

func (s *TeamService) GetTeamsByGame(gameID string) ([]*teamDto.TeamResponse, error) {
	teams, err := s.repo.FindByGameID(gameID)

	var response []*teamDto.TeamResponse

	if err != nil {
		return response, err
	}

	for i := 0; i < len(teams); i++ {
		team := teams[i]

		teamTransformed := s.transformResponse(team)

		response = append(response, teamTransformed)
	}

	return response, nil
}

func (s *TeamService) GetRanking(id string) ([]*rankingDto.RankingResponse, error) {
	rankings, err := s.rankingRepo.FindAllByTeam(id)

	var response []*rankingDto.RankingResponse

	if err != nil {
		return response, err
	}

	for i := 0; i < len(rankings); i++ {
		ranking := rankings[i]

		rankingTransformed := &rankingDto.RankingResponse{
			Id:          ranking.ID,
			GameId:      ranking.GameID,
			TeamId:      ranking.TeamID,
			TotalGames:  ranking.TotalGames,
			TotalWins:   ranking.TotalWins,
			TotalLosses: ranking.TotalLosses,
			WinRate:     ranking.WinRate,
		}

		response = append(response, rankingTransformed)
	}

	return response, nil
}

func (s *TeamService) transformResponse(team *entities.Team) *teamDto.TeamResponse {
	return &teamDto.TeamResponse{
		Id:        team.ID,
		SessionId: team.SessionID,
		GameId:    team.GameID,
	}
}
