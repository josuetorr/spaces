package handlers

import (
	"gitlab.com/josuetorr/spaces/internal/models"
	"gitlab.com/josuetorr/spaces/internal/services"
)

type ActorService interface {
	Create(a services.CreateActorData) error
	Get(by string, value string) (*models.Actor, error)
	GetFollowing(id string) ([]models.Actor, error)
}

const ActivityPubContentType = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
