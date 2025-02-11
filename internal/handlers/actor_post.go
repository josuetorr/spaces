package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"gitlab.com/josuetorr/spaces/internal/models"
	"gitlab.com/josuetorr/spaces/internal/services"
)

type ActorService interface {
	Create(a services.CreateActorData) error
	Get(id string) (*models.Actor, error)
}

type PostActorHandler struct {
	log          *slog.Logger
	actorService ActorService
}

func NewPostActorHandler(log *slog.Logger, actorService services.ActorService) *PostActorHandler {
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

	a, _ := h.actorService.Get(data.Username)
	if a != nil {
		http.Error(w, "User already exists", http.StatusUnprocessableEntity)
		return
	}

	if err := h.actorService.Create(data); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	w.WriteHeader(http.StatusCreated)
}
