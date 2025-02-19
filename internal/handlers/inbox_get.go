package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type InboxService interface {
	GetInboxByActorId(id string) ([]*Activity, error)
}

type GetInboxHandler struct {
	inboxServce InboxService
}

func NewGetInboxHandler(inboxService InboxService) *GetInboxHandler {
	return &GetInboxHandler{inboxServce: inboxService}
}

func (h *GetInboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	as, err := h.inboxServce.GetInboxByActorId(username)
	if err != nil {
		http.Error(w, "Unable to fetch inbox", http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", as)
}
