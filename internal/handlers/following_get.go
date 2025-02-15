package handlers

import (
	"fmt"
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

	a, err := h.actorService.Get("id", username)
	if err != nil {
		if err.Error() == "Invalid query" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if a == nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	following, err := h.actorService.GetFollowing(username)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	fmt.Printf("following: %v, err: %v\n", following, err)

	w.Header().Set("Content-Type", ActivityPubContentType)
}
