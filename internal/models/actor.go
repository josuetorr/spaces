package models

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	ap "github.com/go-ap/activitypub"
	"gitlab.com/josuetorr/spaces/internal/utils"
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
