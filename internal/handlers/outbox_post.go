package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"gitlab.com/josuetorr/spaces/internal/services"
	"gitlab.com/josuetorr/spaces/internal/utils"
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

	createdActivity, err := h.activityService.ActivityGetByDocId(docId)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	instanceName := utils.GetServerURL()
	url := ""
	if strings.Contains(data.Object, instanceName) {
		url = strings.Replace(data.Object, instanceName, fmt.Sprintf("http://localhost:%s", utils.GetServerPort()), 1)
	}

	activityBytes, err := json.Marshal(createdActivity)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	client := http.DefaultClient
	res, err := client.Post(url+"/inbox", ActivityPubContentType, bytes.NewBuffer(activityBytes))
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Unable to contact object's server", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(res.StatusCode)
}
