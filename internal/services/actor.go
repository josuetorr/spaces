package services

import (
	"fmt"
	"os"

	"gitlab.com/josuetorr/spaces/internal/data"
	"gitlab.com/josuetorr/spaces/internal/models"
)

type ActorService struct {
	repo data.ActorRepo
}

func NewActorService(repo data.ActorRepo) ActorService {
	return ActorService{repo: repo}
}

func (s ActorService) Create(data CreateActorData) error {
	domain := os.Getenv("DOMAIN")
	a := models.Actor{
		Id:        fmt.Sprintf("%s/%s", domain, data.Username),
		Type:      models.Person,
		Inbox:     fmt.Sprintf("%s/inbox", domain),
		Outbox:    fmt.Sprintf("%s/outbox", domain),
		Following: fmt.Sprintf("%s/following", domain),
		Followers: fmt.Sprintf("%s/followers", domain),
		Liked:     fmt.Sprintf("%s/liked", domain),
	}

	if err := s.repo.Create(a); err != nil {
		return err
	}
	return nil
}

type CreateActorData struct {
	Username string
	Email    string
	Type     models.ActorType
}
