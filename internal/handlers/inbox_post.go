package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"gitlab.com/josuetorr/spaces/internal/services"
)

type PostInboxHandler struct {
	log             *slog.Logger
	activityService ActivityService
}

func NewPostInboxHandler(log *slog.Logger, activityService ActivityService) *PostInboxHandler {
	return &PostInboxHandler{log: log, activityService: activityService}
}

func (h *PostInboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data services.CreateActivityData
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	fmt.Printf("hello!: %+v\n", data)
	// TODO: finish follow flow:
	// 1. client -> follow req -> follower outbox
	// 2. follower server -> generated follow req -> followee outbox
	// 3. followee server -> generated accept/reject req -> follower inbox
	// 4. follower server sets followee to follower's "followers collection"

	// NOTE: Example
	// alice wants to follow bob: alice -> post to her outbox
	// server -> post bob's inbox
	// bob accept/reject -> post to his outbox
	// server -> post alice's inbox
}
