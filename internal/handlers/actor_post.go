package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"gitlab.com/josuetorr/spaces/internal/services"
)

type PostActorHandler struct {
	log          *slog.Logger
	actorService ActorService
}

func NewPostActorHandler(log *slog.Logger, actorService ActorService) *PostActorHandler {
	return &PostActorHandler{log: log, actorService: actorService}
}

func (h *PostActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data services.CreateActorData
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	if err := data.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	a, _ := h.actorService.Get("id", data.Username)
	if a != nil {
		http.Error(w, "User already exists", http.StatusUnprocessableEntity)
		return
	}

	a, _ = h.actorService.Get("email", data.Email)
	if a != nil {
		http.Error(w, "User already exists", http.StatusUnprocessableEntity)
		return
	}

	if err := h.actorService.Create(data); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`)
	w.WriteHeader(http.StatusCreated)
}
