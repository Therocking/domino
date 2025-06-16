package services

import (
	"errors"
	dto "githup/Therocking/dominoes/internal/dtos/game"
	"githup/Therocking/dominoes/internal/entities"
	"githup/Therocking/dominoes/internal/repositories"

	"github.com/google/uuid"
)

type GameService interface {
	AddPoint(dto *dto.CreateGame) error
	GetPointsByGameId(gameId string) ([]*entities.GamePoint, error)
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

func (s *gameService) AddPoint(dto *dto.CreateGame) error {
	game, err := s.gameRepo.FindByID(dto.GameId)
	if err != nil {
		return err
	}

	if game.Completed {
		return errors.New("game already completed")
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

	s.completeGame(game, dto.TeamId)

	return createErr
}

func (s *gameService) GetPointsByGameId(gameId string) ([]*entities.GamePoint, error) {
	gamePoint, err := s.gamePointRepo.FindByGameID(gameId)

	if err != nil {
		return nil, err
	}

	return gamePoint, err
}

func (s *gameService) completeGame(game *entities.Game, teamId string) error {
	isGreaten, err := s.isTotalPointGratenThenGamePoint(game.ID, teamId)

	if err != nil {
		return err
	}

	if isGreaten {
		game.Completed = true
		if err := s.gameRepo.Update(game); err != nil {
			return err
		}
	}

	return nil
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
