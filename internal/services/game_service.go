package services

import (
	"errors"
	dto "githup/Therocking/dominoes/internal/dtos/game"
	gameDto "githup/Therocking/dominoes/internal/dtos/game"
	"githup/Therocking/dominoes/internal/entities"
	"githup/Therocking/dominoes/internal/repositories"

	"github.com/google/uuid"
)

type GameService interface {
	AddPoint(dto *gameDto.CreateGamePoint) (*gameDto.GameCompletedResponse, error)
	GetPointsByGameId(gameId string) ([]*dto.GamePointResponse, error)
}

type gameService struct {
	gameRepo      repositories.GameRepository
	gamePointRepo repositories.GamePointRepository
	rankingRepo   repositories.RankingRepository
}

func NewGameService(
	gameRepo repositories.GameRepository,
	gamePointRepo repositories.GamePointRepository,
	rankingRepo repositories.RankingRepository,
) GameService {
	return &gameService{
		gameRepo:      gameRepo,
		gamePointRepo: gamePointRepo,
		rankingRepo:   rankingRepo,
	}
}

func (s *gameService) AddPoint(dto *gameDto.CreateGamePoint) (*gameDto.GameCompletedResponse, error) {
	game, err := s.gameRepo.FindByID(dto.GameId)
	if err != nil {
		return nil, err
	}

	if game.Completed {
		return nil, errors.New("game already completed")
	}

	gamePoint := &entities.GamePoint{
		Base: entities.Base{
			ID: uuid.New().String(),
		},
		GameID: dto.GameId,
		Point:  dto.Point,
		TeamID: dto.TeamId,
	}

	createErr := s.gamePointRepo.Create(gamePoint)

	if createErr != nil {
		return nil, err
	}

	response, err := s.completeGame(game, dto.TeamId)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *gameService) GetPointsByGameId(gameId string) ([]*dto.GamePointResponse, error) {
	gamePoints, err := s.gamePointRepo.FindByGameID(gameId)

	if err != nil {
		return nil, err
	}

	var response []*dto.GamePointResponse

	for i := 0; i < len(gamePoints); i++ {
		gamePoint := gamePoints[i]

		gamePointTransformed := dto.GamePointResponse{
			Id:     gamePoint.ID,
			GameId: gamePoint.GameID,
			TeamId: gamePoint.TeamID,
		}

		response = append(response, &gamePointTransformed)
	}

	return response, nil
}

func (s *gameService) completeGame(game *entities.Game, teamId string) (*gameDto.GameCompletedResponse, error) {
	isGreaten, err := s.isTotalPointGratenThenGamePoint(game.ID, teamId)

	if err != nil {
		return nil, err
	}

	if !isGreaten {
		return nil, nil
	}

	game.Completed = true
	if err := s.gameRepo.Update(game); err != nil {
		return nil, err
	}

	response := &gameDto.GameCompletedResponse{
		Message:      "Game ended",
		WinnerTeamId: teamId,
	}

	return response, nil
}

func (s *gameService) isTotalPointGratenThenGamePoint(gameId, teamId string) (bool, error) {
	points, err := s.gamePointRepo.FindAllByGameAndTeamId(gameId, teamId)

	if err != nil {
		return false, err
	}

	var totalPoints int
	for _, point := range points {
		totalPoints += point.Point
	}

	return totalPoints >= 200, nil
}
