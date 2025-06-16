package services

import (
	"githup/Therocking/dominoes/internal/entities"
	"githup/Therocking/dominoes/internal/repositories"

	"github.com/google/uuid"
)

type SessionService interface {
	CreateSession(deviceID string) (*entities.Session, error)
}

type sessionService struct {
	sessionRepo repositories.SessionRepository
	gameRepo    repositories.GameRepository
	teamRepo    repositories.TeamRepository
}

func NewSessionService(
	sessionRepo repositories.SessionRepository,
	gameRepo repositories.GameRepository,
	teamRepo repositories.TeamRepository,
) SessionService {
	return &sessionService{
		sessionRepo: sessionRepo,
		gameRepo:    gameRepo,
		teamRepo:    teamRepo,
	}
}

func (s *sessionService) CreateSession(deviceID string) (*entities.Session, error) {
	session := &entities.Session{
		Base: entities.Base{
			ID: uuid.New().String(),
		},
		DeviceID: deviceID,
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}

	game := &entities.Game{
		Base: entities.Base{
			ID: uuid.New().String(),
		},
	}
	if err := s.gameRepo.Create(game); err != nil {
		return nil, err
	}

	teams := []*entities.Team{
		{
			Base: entities.Base{
				ID: uuid.New().String(),
			},
			Name:      "team1",
			SessionID: session.ID,
			GameID:    game.ID,
		},
		{
			Base: entities.Base{
				ID: uuid.New().String(),
			},
			Name:      "team2",
			SessionID: session.ID,
			GameID:    game.ID,
		},
	}

	for _, team := range teams {
		if err := s.teamRepo.Create(team); err != nil {
			return nil, err
		}
	}

	return session, nil
}
