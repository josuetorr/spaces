package handlers

import (
	"log/slog"
	"net/http"
)

type PostInboxHandler struct {
	log *slog.Logger
}

func NewPostInboxHandler(log *slog.Logger) *PostInboxHandler {
	return &PostInboxHandler{log: log}
}

func (h *PostInboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
