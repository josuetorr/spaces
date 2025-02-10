package services

import "gitlab.com/josuetorr/spaces/internal/data"

type ActorService struct {
	repo data.ActorRepo
}

func NewActorService(repo data.ActorRepo) ActorService {
	return ActorService{repo: repo}
}
