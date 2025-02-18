package handlers

import (
	"log/slog"
	"net/http"
)

type PostInboxHandler struct {
	log             *slog.Logger
	activityService ActivityService
}

func NewPostInboxHandler(log *slog.Logger, activityService ActivityService) *PostInboxHandler {
	return &PostInboxHandler{log: log, activityService: activityService}
}

func (h *PostInboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
