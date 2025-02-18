package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"gitlab.com/josuetorr/spaces/internal/services"
)

type PostOutboxHandler struct {
	log             *slog.Logger
	activityService ActivityService
}

func NewPostOutboxHandler(log *slog.Logger, activityService ActivityService) *PostOutboxHandler {
	return &PostOutboxHandler{log: log, activityService: activityService}
}

func (h *PostOutboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data services.CreateActivityData
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	docId, err := h.activityService.ActivityCreate(data)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	createdActivity, err := h.activityService.ActivityGetById(docId)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// client := http.DefaultClient
	println(createdActivity)
}
