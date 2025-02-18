package services

import (
	"fmt"

	ap "github.com/go-ap/activitypub"
	"gitlab.com/josuetorr/spaces/internal/models"
	"gitlab.com/josuetorr/spaces/internal/utils"
)

type CreateActorData struct {
	Type              models.ActorType
	Firstname         string
	Lastname          string
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

type ActorRepo interface {
	Create(data *models.Actor) error

	Exists(id string) (bool, error)
	GetById(id string) (*models.Actor, error)
	GetByEmail(email string) (*models.Actor, error)
	GetFollowing(id string) ([]models.Actor, error)
}

type ActorService struct {
	repo ActorRepo
}

func NewActorService(repo ActorRepo) ActorService {
	return ActorService{repo: repo}
}

func (s ActorService) Create(data CreateActorData) error {
	id := utils.GetFullId("users", data.Username)
	a := &models.Actor{
		Id:                id,
		Type:              data.Type,
		Firstname:         data.Firstname,
		Lastname:          data.Lastname,
		PreferredUsername: data.PreferredUsername,
		Email:             data.Email,
		Follows:           []models.Actor{},
	}

	if data.PreferredUsername == "" {
		a.PreferredUsername = data.Username
	}

	if err := s.repo.Create(a); err != nil {
		return err
	}
	return nil
}

func (s ActorService) Exists(id string) (bool, error) {
	id = utils.GetFullId("users", id)
	return s.repo.Exists(id)
}

func (s ActorService) GetById(id string) (*models.Actor, error) {
	id = utils.GetFullId("users", id)
	return s.repo.GetById(id)
}

func (s ActorService) GetByEmail(email string) (*models.Actor, error) {
	return s.repo.GetByEmail(email)
}

// NOTE: for now, the following collection is not ordered not paginated
//
//	{
//	  "@context": "https://www.w3.org/ns/activitystreams",
//	  "summary": "Sally's notes",
//	  "type": "Collection",
//	  "totalItems": 2,
//	  "items": [
//	    {
//	      "type": "Note",
//	      "name": "A Simple Note"
//	    },
//	    {
//	      "type": "Note",
//	      "name": "Another Simple Note"
//	    }
//	  ]
//	}
func (s ActorService) GetFollowing(id string) (*ap.Collection, error) {
	userId := utils.GetFullId("users", id)

	following, err := s.repo.GetFollowing(id)
	followingLen := uint(len(following))

	collectionID := userId + "/following"
	c := ap.CollectionNew(ap.ID(collectionID))
	c.TotalItems = followingLen
	items := make(ap.ItemCollection, followingLen)
	for i, f := range following {
		items[i] = ap.ID(f.Id)
	}
	c.Items = items

	return c, err
}
