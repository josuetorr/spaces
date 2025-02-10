package handlers

import (
	"net/http"

	"gitlab.com/josuetorr/spaces/internal/models"
)

var store = make(map[string]models.Actor)

type WebFingerHandler struct{}

func NewWebFingerHandler() *WebFingerHandler {
	return &WebFingerHandler{}
}

func (h *WebFingerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
