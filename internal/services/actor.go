package services

import (
	"gitlab.com/josuetorr/spaces/internal/data"
	"gitlab.com/josuetorr/spaces/internal/models"
)

type CreateActorData struct {
	Username string
	Email    string
	Type     models.ActorType
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
		Id:   data.Username,
		Type: models.Person,
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
