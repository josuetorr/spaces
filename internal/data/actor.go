package data

import (
	"log/slog"

	"github.com/go-kivik/kivik/v4"
	"gitlab.com/josuetorr/spaces/internal/models"
)

type ActorRepo struct {
	db  *kivik.Client
	log *slog.Logger
}

func NewActorRepo(db *kivik.Client, log *slog.Logger) ActorRepo {
	return ActorRepo{db: db, log: log}
}

func (r ActorRepo) Create(a *models.Actor) error {
	panic("implement create actor")
}

func (r ActorRepo) Get(by string, value string) (*models.Actor, error) {
	panic("implement get actor")
}

func (r ActorRepo) GetFollowing(id string) ([]models.Actor, error) {
	panic("implement get following")
}
