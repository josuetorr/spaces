package handlers

import (
	ap "github.com/go-ap/activitypub"
	"gitlab.com/josuetorr/spaces/internal/models"
	"gitlab.com/josuetorr/spaces/internal/services"
)

type ActorService interface {
	Create(a services.CreateActorData) error
	Exists(id string) (bool, error)
	GetById(id string) (*models.Actor, error)
	GetByEmail(email string) (*models.Actor, error)
	GetFollowing(id string) (*ap.Collection, error)
}

const ActivityPubContentType = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
