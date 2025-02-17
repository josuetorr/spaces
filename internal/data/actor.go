package data

import (
	"context"
	"log/slog"

	"github.com/go-kivik/kivik/v4"
	"gitlab.com/josuetorr/spaces/internal/models"
)

type ActorRepo struct {
	db  *kivik.DB
	log *slog.Logger
}

func NewActorRepo(db *kivik.DB, log *slog.Logger) ActorRepo {
	return ActorRepo{db: db, log: log}
}

func (r ActorRepo) Create(a *models.Actor) error {
	_, err := r.db.Put(context.TODO(), a.Id, a)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}

func (r ActorRepo) GetById(id string) (*models.Actor, error) {
	var a models.Actor
	if err := r.db.Get(context.TODO(), id).ScanDoc(&a); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r ActorRepo) GetByEmail(email string) (*models.Actor, error) {
	var a models.Actor
	if err := r.db.Query(context.TODO(), "_design/users", "_views/by-email").ScanDoc(&a); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r ActorRepo) GetFollowing(id string) ([]models.Actor, error) {
	return nil, nil
}
