package services

import (
	sessionDto "githup/Therocking/dominoes/internal/dtos/sessions"
	dto "githup/Therocking/dominoes/internal/dtos/team"
	"githup/Therocking/dominoes/internal/entities"
	"githup/Therocking/dominoes/internal/repositories"

	"github.com/google/uuid"
)

type SessionService interface {
	CreateSession(deviceID string) (*sessionDto.SessionCreatedResponse, error)
	GetByDeviceId(deviceID string) (*sessionDto.SessionResponse, error)
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

func (s *sessionService) CreateSession(deviceID string) (*sessionDto.SessionCreatedResponse, error) {
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

	return &sessionDto.SessionCreatedResponse{
		SessionId: session.ID,
	}, nil
}

func (s *sessionService) GetByDeviceId(deviceID string) (*sessionDto.SessionResponse, error) {
	session, err := s.sessionRepo.FindByDeviceID(deviceID)

	if err != nil {
		return nil, err
	}

	var teams []*dto.TeamResponse
	for i := 0; i < len(session.Teams); i++ {
		team := session.Teams[i]

		teams = append(teams, &dto.TeamResponse{
			Id:        team.ID,
			SessionId: team.SessionID,
			GameId:    team.GameID,
		})
	}

	response := &sessionDto.SessionResponse{
		Id:    session.ID,
		Teams: teams,
	}

	return response, nil
}
