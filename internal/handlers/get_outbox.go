package handlers

import (
	"net/http"
)

type GetOutboxHandler struct{}

func NewGetOutboxHandler() *GetOutboxHandler {
	return &GetOutboxHandler{}
}

func (h *GetOutboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
