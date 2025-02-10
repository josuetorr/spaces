package data

import (
	"database/sql"

	"gitlab.com/josuetorr/spaces/internal/models"
)

const (
	getActorQuery    = "SELECT * FROM actors WHERE id=?;"
	insertActorQuery = "INSERT INTO actors(id, type, inbox, outbox, following, followers, liked) VALUES(?,?,?,?,?,?,?)"
)

type ActorRepo struct {
	db *sql.DB
}

func NewActorRepo(db *sql.DB) ActorRepo {
	return ActorRepo{db: db}
}

func (r ActorRepo) Get(id string) (*models.Actor, error) {
	stmt, err := r.db.Prepare(getActorQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var actor models.Actor
	result := stmt.QueryRow(id)
	if err := result.Scan(
		&actor.Id,
		&actor.Type,
		&actor.Inbox,
		&actor.Outbox,
		&actor.Following,
		&actor.Followers,
		&actor.Liked); err != nil {
		return nil, err
	}

	return &actor, nil
}

func (r ActorRepo) Create(a models.Actor) error {
	stmt, err := r.db.Prepare(insertActorQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Id, a.Type, a.Id, a.Outbox, a.Following, a.Followers, a.Liked)

	return err
}
