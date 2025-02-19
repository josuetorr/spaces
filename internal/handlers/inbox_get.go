package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/josuetorr/spaces/internal/utils"
)

type InboxService interface {
	GetInboxByActorId(id string) ([]*Activity, error)
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

	as, err := h.inboxServce.GetInboxByActorId(id)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Unable to fetch inbox", http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", as)
}
