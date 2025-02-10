package handlers

import (
	"net/http"

	"gitlab.com/josuetorr/spaces/internal/services"
)

type PostActorHandler struct {
	actorService services.ActorService
}

func NewPostActorHandler(actorService services.ActorService) *PostActorHandler {
	return &PostActorHandler{actorService: actorService}
}

func (h *PostActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
