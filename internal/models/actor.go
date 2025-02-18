package models

import (
	ap "github.com/go-ap/activitypub"
)

type ActorType = ap.ActivityVocabularyType

const (
	Application  ActorType = ap.ApplicationType
	Group        ActorType = ap.GroupType
	Organization ActorType = ap.OrganizationType
	Person       ActorType = ap.PersonType
	Service      ActorType = ap.ServiceType
)

type Actor struct {
	Uid               string    `json:"uid,omitempty"`
	Id                string    `json:"id,omitempty"`
	Type              ActorType `json:"type,omitempty"`
	Firstname         string    `json:"firstname, omitempty"`
	Lastname          string    `json:"lastname, omitempty"`
	PreferredUsername string    `json:"preferredUsername,omitempty"`
	Email             string    `json:"email,omitempty"`
	Follows           []Actor   `json:"follows,omitempty"`
}

func (a Actor) JSON() *ap.Actor {
	apActor := ap.ActorNew(ap.ID(a.Id), ap.ActivityVocabularyType(a.Type))
	apActor.Inbox = ap.IRI(a.Id + "/inbox")
	apActor.Outbox = ap.IRI(a.Id + "/outbox")
	apActor.Following = ap.IRI(a.Id + "/following")
	apActor.Followers = ap.IRI(a.Id + "/followers")

	return apActor
}
