package data

import (
	"database/sql"

	"gitlab.com/josuetorr/spaces/internal/models"
)

const (
	getActorQuery    = "SELECT * FROM actors WHERE id=?;"
	insertActorQuery = "INSERT INTO actors(id, type, preferredUsername) VALUES(?,?,?)"
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
		&actor.PreferredUsername); err != nil {
		return nil, err
	}

	return &actor, nil
}

func (r ActorRepo) Create(a *models.Actor) error {
	stmt, err := r.db.Prepare(insertActorQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Id, a.Type, a.PreferredUsername)

	return err
}
