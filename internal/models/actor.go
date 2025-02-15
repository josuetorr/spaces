package models

import (
	"fmt"
	"os"
	"reflect"
	"strings"

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

func (a Actor) ToDto() *ap.Actor {
	id := fmt.Sprintf("https://%s/%s", os.Getenv("SPACES_SERVER_NAME"), a.Id)
	apActor := ap.ActorNew(ap.ID(id), ap.ActivityVocabularyType(a.Type))
	apActor.Inbox = ap.IRI(id + "/inbox")
	apActor.Outbox = ap.IRI(id + "/outbox")
	apActor.Following = ap.IRI(id + "/following")
	apActor.Followers = ap.IRI(id + "/followers")
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
