package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/dgraph-io/dgo/v240"
	"github.com/dgraph-io/dgo/v240/protos/api"
	"gitlab.com/josuetorr/spaces/internal/models"
)

const (
	getActorByIdQuery = `
  query{
      q(func: type(Actor)) @filter(eq(%s, %s)){
        id
        type
        firstname
        lastname
        preferredUsername
        email 
      }
    }
  `
	getActorFollowingByIdQuery = `
  query following($id: string){
    q(func: eq(id, $id)){
      follows {
        id
      }
    }
  }
  `
)

type ActorRepo struct {
	dg  *dgo.Dgraph
	log *slog.Logger
}

func NewActorRepo(dg *dgo.Dgraph, log *slog.Logger) ActorRepo {
	return ActorRepo{dg: dg, log: log}
}

func (r ActorRepo) Get(by string, value string) (*models.Actor, error) {
	allowedFields := map[string]bool{
		"id":    true,
		"email": true,
	}

	if !allowedFields[by] {
		return nil, errors.New("Invalid query")
	}

	txn := r.dg.NewTxn()
	// TODO: pass ctx
	res, err := txn.Query(context.Background(), fmt.Sprintf(getActorByIdQuery, by, value))
	if err != nil {
		return nil, err
	}

	type Root struct {
		Q []models.Actor `json:"q"`
	}
	var root Root
	if err := json.Unmarshal(res.Json, &root); err != nil {
		return nil, err
	}

	if len(root.Q) == 0 {
		return nil, nil
	}

	return &root.Q[0], nil
}

func (r ActorRepo) GetFollowing(id string) ([]models.Actor, error) {
	vars := make(map[string]string)
	vars["$id"] = id

	txn := r.dg.NewTxn()
	res, err := txn.QueryWithVars(context.Background(), getActorFollowingByIdQuery, vars)
	if err != nil {
		return nil, err
	}

	type Root struct {
		Q []models.Actor `json:"q"`
	}
	var root Root
	if err := json.Unmarshal(res.Json, &root); err != nil {
		return nil, err
	}

	if len(root.Q) == 0 {
		return nil, err
	}

	return root.Q[0].Follows, nil
}

func (r ActorRepo) Create(a *models.Actor) error {
	nquads := a.NQuads()
	mut := &api.Mutation{CommitNow: true, SetNquads: nquads}
	_, err := r.dg.NewTxn().Mutate(context.Background(), mut)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}

	return nil
}
