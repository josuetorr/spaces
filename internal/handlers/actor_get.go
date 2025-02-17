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

	a, err := h.actorService.GetById(username)
	if err != nil {
		if err.Error() == "Not Found: missing" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if a == nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", ActivityPubContentType)
	json.NewEncoder(w).Encode(a.JSON())
}
