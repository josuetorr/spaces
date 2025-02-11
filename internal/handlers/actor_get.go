package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type GetActorHandler struct {
	actorService ActorService
}

func NewGetActorHandler(ActorService ActorService) *GetActorHandler {
	return &GetActorHandler{actorService: ActorService}
}

func (h *GetActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	a, _ := h.actorService.Get(username)

	if a == nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	json.NewEncoder(w).Encode(a.ToDto())
}
