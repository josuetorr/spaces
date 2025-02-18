package handlers

import (
	"gitlab.com/josuetorr/spaces/internal/services"
)

type (
	Actor      = services.Actor
	Collection = services.Collection
)

type ActorService interface {
	Create(a services.CreateActorData) error
	Exists(id string) (bool, error)
	GetById(id string) (*Actor, error)
	GetByEmail(email string) (*Actor, error)
	GetFollowing(id string) (*Collection, error)
}

const ActivityPubContentType = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
