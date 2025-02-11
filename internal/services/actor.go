package services

import (
	"fmt"

	"gitlab.com/josuetorr/spaces/internal/data"
	"gitlab.com/josuetorr/spaces/internal/models"
)

type CreateActorData struct {
	Username          string
	PreferredUsername string
	Email             string
	Type              models.ActorType
}

func (data CreateActorData) Validate() error {
	if data.Username == "" {
		return fmt.Errorf("Must provided username")
	}
	if data.Type == "" {
		return fmt.Errorf("Must provided actor type")
	}
	// TODO: validate email format
	if data.Email == "" {
		return fmt.Errorf("Must provided email")
	}
	return nil
}

type ActorRepo interface {
	Get(id string) (*models.Actor, error)
	Create(data *models.Actor) error
}

type ActorService struct {
	repo ActorRepo
}

func NewActorService(repo data.ActorRepo) ActorService {
	return ActorService{repo: repo}
}

func (s ActorService) Create(data CreateActorData) error {
	a := &models.Actor{
		Id:    data.Username,
		Type:  data.Type,
		Email: data.Email,
	}

	if data.PreferredUsername == "" {
		a.PreferredUsername = data.Username
	}

	if err := s.repo.Create(a); err != nil {
		return err
	}
	return nil
}

func (s ActorService) Get(id string) (*models.Actor, error) {
	a, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return a, nil
}
