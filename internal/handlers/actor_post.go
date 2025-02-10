package handlers

import (
	"net/http"

	"gitlab.com/josuetorr/spaces/internal/data"
)

type PostActorHandler struct {
	repo *data.ActorRepo
}

func NewPostActorHandler(repo *data.ActorRepo) *PostActorHandler {
	return &PostActorHandler{repo: repo}
}

func (h *PostActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
