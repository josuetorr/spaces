package models

import (
	"fmt"
	"os"
	"reflect"
	"strings"
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
	Uid               string    `json:"uid,omitempty"`
	Id                string    `json:"id,omitempty"`
	Type              ActorType `json:"type,omitempty"`
	PreferredUsername string    `json:"preferredUsername,omitempty"`
	Email             string    `json:"email,omitempty"`
	Follows           []Actor   `json:"follows,omitempty"`
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

func (a Actor) NQuads() []byte {
	format := "_:%s <%s> \"%s\" .\n"
	nquads := fmt.Sprintf(format, a.Id, "dgraph.type", "Actor")

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	for i := 0; i < t.NumField(); i++ {
		field := strings.ToLower(t.Field(i).Name)
		fieldValue := v.Field(i)

		if field == "uid" || fieldValue.IsZero() {
			continue
		}

		nquads = nquads + fmt.Sprintf(format, a.Id, field, fieldValue)
	}

	return []byte(nquads)
}
