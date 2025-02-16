package services

import (
	"fmt"
	"os"

	ap "github.com/go-ap/activitypub"
	"gitlab.com/josuetorr/spaces/internal/data"
	"gitlab.com/josuetorr/spaces/internal/models"
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
	Get(by string, value string) (*models.Actor, error)
	Create(data *models.Actor) error
	GetFollowing(id string) ([]models.Actor, error)
}

type ActorService struct {
	repo ActorRepo
}

func NewActorService(repo data.ActorRepo) ActorService {
	return ActorService{repo: repo}
}

func (s ActorService) Create(data CreateActorData) error {
	a := &models.Actor{
		Id:                data.Username,
		Type:              data.Type,
		Firstname:         data.Firstname,
		Lastname:          data.Lastname,
		PreferredUsername: data.PreferredUsername,
		Email:             data.Email,
	}

	if data.PreferredUsername == "" {
		a.PreferredUsername = data.Username
	}

	if err := s.repo.Create(a); err != nil {
		return err
	}
	return nil
}

func (s ActorService) Get(by string, value string) (*models.Actor, error) {
	return s.repo.Get(by, value)
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
	// TODO: remove this bullcrap and simply store the ids in full
	serverName := fmt.Sprintf("https://%s", os.Getenv("SPACES_SERVER_NAME"))
	userId := fmt.Sprintf("%s/users/%s", serverName, id)

	following, err := s.repo.GetFollowing(id)
	followingLen := uint(len(following))

	collectionID := userId + "/following"
	c := ap.CollectionNew(ap.ID(collectionID))
	c.TotalItems = followingLen
	items := make(ap.ItemCollection, followingLen)
	for i, f := range following {
		items[i] = ap.ID(fmt.Sprintf("%s/users/%s", serverName, f.Id))
	}
	c.Items = items

	return c, err
}
