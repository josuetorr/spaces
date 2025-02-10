package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"gitlab.com/josuetorr/spaces/internal/models"
	"gitlab.com/josuetorr/spaces/internal/services"
)

type ActorService interface {
	Create(a *models.Actor) error
}

type PostActorHandler struct {
	log          *slog.Logger
	actorService ActorService
}

type createActorRequest struct {
	username string
	email    string
}

func NewPostActorHandler(log *slog.Logger, actorService services.ActorService) *PostActorHandler {
	return &PostActorHandler{log: log, actorService: actorService}
}

func (h *PostActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data createActorRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	domain := os.Getenv("domain")
	actor := models.Actor{
		Id:        fmt.Sprintf("%s/%s", domain, data.username),
		Type:      models.Person,
		Inbox:     fmt.Sprintf("%s/inbox", domain),
		Outbox:    fmt.Sprintf("%s/outbox", domain),
		Following: fmt.Sprintf("%s/following", domain),
		Followers: fmt.Sprintf("%s/followers", domain),
		Liked:     fmt.Sprintf("%s/liked", domain),
	}

	if err := h.actorService.Create(&actor); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/activity+json")
	w.WriteHeader(http.StatusCreated)
}
