package services

import (
	"gitlab.com/josuetorr/spaces/internal/data"
	"gitlab.com/josuetorr/spaces/internal/models"
)

type ActorService struct {
	repo data.ActorRepo
}

func NewActorService(repo data.ActorRepo) ActorService {
	return ActorService{repo: repo}
}

func (s ActorService) Create(a *models.Actor) error {
	if err := s.repo.Create(*a); err != nil {
		return err
	}
	return nil
}
