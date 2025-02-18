package handlers

import (
	"gitlab.com/josuetorr/spaces/internal/services"
)

type (
	Actor      = services.Actor
	Activity   = services.Activity
	Collection = services.Collection
)

type ActorService interface {
	ActorCreate(a services.CreateActorData) (string, error)
	ActorExists(id string) (bool, error)
	ActorGetById(id string) (*Actor, error)
	ActorGetByEmail(email string) (*Actor, error)
	ActorGetFollowing(id string) (*Collection, error)
}

type ActivityService interface {
	ActivityCreate(a services.CreateActivityData) (string, error)
	ActivityExists(id string) (bool, error)
	ActivityGetById(id string) (*Activity, error)
	ActivityGetByDocId(id string) (*Activity, error)
}

const ActivityPubContentType = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
