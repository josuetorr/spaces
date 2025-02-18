package services

import (
	"fmt"

	ap "github.com/go-ap/activitypub"
	"gitlab.com/josuetorr/spaces/internal/utils"
)

type CreateActorData struct {
	Type              string
	Name              string
	Username          string
	PreferredUsername string
	Email             string
}

func (data CreateActorData) Validate() error {
	if data.Username == "" {
		return fmt.Errorf("Must provided username")
	}
	if data.Type == "" {
		return fmt.Errorf("Must provided actor type")
	}
	// TODO: validate email format
	if data.Email == "" {
		return fmt.Errorf("Must provided email")
	}
	return nil
}

type (
	Actor           = ap.Actor
	Collection      = ap.Collection
	ActorRepository interface {
		Repository[Actor]
		GetByEmail(string) (*Actor, error)
		GetFollowing(string) (ap.IRIs, error)
	}
	ActorService struct {
		repo ActorRepository
	}
)

func NewActorService(repo ActorRepository) ActorService {
	return ActorService{repo: repo}
}

func (s ActorService) ActorCreate(data CreateActorData) (string, error) {
	id := utils.GetFullId("users", data.Username)
	a := ap.ActorNew(ap.ID(id), ap.ActivityVocabularyType(data.Type))

	preferredUsername := data.PreferredUsername
	if preferredUsername == "" {
		preferredUsername = data.Username
	}

	a.Name = ap.NaturalLanguageValuesNew(ap.LangRefValueNew("en", data.Name))
	a.PreferredUsername = ap.NaturalLanguageValuesNew(ap.LangRefValueNew("en", preferredUsername))
	a.Inbox = ap.IRI(id + "/inbox")
	a.Outbox = ap.IRI(id + "/outbox")
	a.Following = ap.IRI(id + "/following")
	a.Followers = ap.IRI(id + "/followers")

	docId, err := s.repo.Create(a)
	if err != nil {
		return docId, err
	}
	return docId, nil
}

func (s ActorService) ActorExists(id string) (bool, error) {
	id = utils.GetFullId("users", id)
	return s.repo.Exists(id)
}

func (s ActorService) ActorGetById(id string) (*Actor, error) {
	id = utils.GetFullId("users", id)
	return s.repo.GetById(id)
}

func (s ActorService) ActorGetByEmail(email string) (*Actor, error) {
	return s.repo.GetByEmail(email)
}

func (s ActorService) ActorGetFollowing(id string) (*Collection, error) {
	userId := utils.GetFullId("users", id)

	following, err := s.repo.GetFollowing(id)
	followingLen := uint(len(following))

	collectionID := userId + "/following"
	c := ap.CollectionNew(ap.ID(collectionID))
	c.TotalItems = followingLen
	items := make(ap.ItemCollection, followingLen)
	for i, f := range following {
		items[i] = f.GetID()
	}
	c.Items = items

	return c, err
}
