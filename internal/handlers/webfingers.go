package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type WebFingerHandler struct {
	log          *slog.Logger
	actorService ActorService
}

func NewWebFingerHandler(log *slog.Logger, actorService ActorService) *WebFingerHandler {
	return &WebFingerHandler{log: log, actorService: actorService}
}

func (h *WebFingerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("resource")

	if query == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	acct := strings.TrimSpace(strings.TrimPrefix(query, "acct:"))

	parts := strings.Split(acct, "@")
	username := parts[0]
	domain := parts[1]
	if domain != os.Getenv("SERVER_NAME") {
		http.Error(w, "Domain not handled by server", http.StatusUnprocessableEntity)
		return
	}

	a, _ := h.actorService.Get(username)
	if a == nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	resp := WebFingerResponse{
		Subject: query,
		Links: []Link{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: fmt.Sprintf("https://%s/%s", domain, username),
			},
		},
	}

	w.Header().Set("Content-Type", "application/activity+json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

type WebFingerResponse struct {
	Subject string `json:"subject"`
	Links   []Link `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}
