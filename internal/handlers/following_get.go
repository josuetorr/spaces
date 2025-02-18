package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type GetFollowingHandler struct {
	log          *slog.Logger
	actorService ActorService
}

func NewGetFollowingHandler(log *slog.Logger, actorService ActorService) *GetFollowingHandler {
	return &GetFollowingHandler{log: log, actorService: actorService}
}

func (h *GetFollowingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	exists, err := h.actorService.Exists(username)

	if !exists {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	following, err := h.actorService.GetFollowing(username)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(following)
	w.Header().Set("Content-Type", ActivityPubContentType)
}
