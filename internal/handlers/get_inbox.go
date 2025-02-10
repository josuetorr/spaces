package handlers

import (
	"net/http"
)

type GetInboxHandler struct{}

func NewGetInboxHandler() *GetInboxHandler {
	return &GetInboxHandler{}
}

func (h *GetInboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
