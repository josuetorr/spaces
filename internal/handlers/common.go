package handlers

import (
	"gitlab.com/josuetorr/spaces/internal/models"
	"gitlab.com/josuetorr/spaces/internal/services"
)

type ActorService interface {
	Create(a services.CreateActorData) error
	Get(by string, value string) (*models.Actor, error)
}
