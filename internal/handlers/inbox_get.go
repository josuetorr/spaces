package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/josuetorr/spaces/internal/utils"
)

type InboxService interface {
	GetInboxByActorId(id string) (*Collection, error)
}

type GetInboxHandler struct {
	log         *slog.Logger
	inboxServce InboxService
}

func NewGetInboxHandler(log *slog.Logger, inboxService InboxService) *GetInboxHandler {
	return &GetInboxHandler{log: log, inboxServce: inboxService}
}

func (h *GetInboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	id := utils.GetFullId("users", username)

	activityCollection, err := h.inboxServce.GetInboxByActorId(id)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Unable to fetch inbox", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", ActivityPubContentType)
	json.NewEncoder(w).Encode(activityCollection)
}
