package models

import (
	"fmt"
	"os"
)

type ActorType string

const (
	Application  ActorType = "Application"
	Group        ActorType = "Group"
	Orginization ActorType = "Orginization"
	Person       ActorType = "Person"
	Service      ActorType = "Service"
)

type Actor struct {
	Id                string    `json:"id"`
	Type              ActorType `json:"type"`
	PreferredUsername string    `json:"preferredUsername"`
	Email             string    `json:"email"`
}

type ActorDto struct {
	Actor
	Context   string `json:"@context"`
	Inbox     string `json:"inbox"`
	Outbox    string `json:"outbox"`
	Following string `json:"following"`
	Followers string `json:"followers"`
	Liked     string `json:"liked"`
}

func (a Actor) ToDto() ActorDto {
	id := fmt.Sprintf("https://%s/%s", os.Getenv("SERVER_NAME"), a.Id)
	return ActorDto{
		Context: "https://www.w3.org/ns/activitystreams",
		Actor: Actor{
			Id:                id,
			Email:             a.Email,
			PreferredUsername: a.PreferredUsername,
			Type:              a.Type,
		},
		Inbox:     id + "/inbox",
		Outbox:    id + "/outbox",
		Following: id + "/following",
		Followers: id + "/followers",
		Liked:     id + "/liked",
	}
}
