package data

import (
	"context"
	"log/slog"

	ap "github.com/go-ap/activitypub"
	"github.com/go-kivik/kivik/v4"
	"gitlab.com/josuetorr/spaces/internal/services"
)

type Actor = services.Actor

type ActorRepo struct {
	Repository[Actor]
}

func NewActorRepo(db *kivik.DB, log *slog.Logger) ActorRepo {
	return ActorRepo{Repository[ap.Actor]{db: db, log: log}}
}

func (r ActorRepo) GetByEmail(email string) (*Actor, error) {
	var a *Actor
	query := map[string]any{
		"selector": map[string]string{
			"email": email,
		},
		"limit": 1,
	}
	rows := r.db.Find(context.TODO(), query)
	defer rows.Close()

	for rows.Next() {
		if err := rows.ScanDoc(&a); err != nil {
			return nil, err
		}
	}
	return a, nil
}

func (r ActorRepo) GetFollowing(id string) (ap.IRIs, error) {
	return ap.IRIs{}, nil
}
