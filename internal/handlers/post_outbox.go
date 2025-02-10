package handlers

import (
	"net/http"
)

type PostOutboxHandler struct{}

func NewPostOutboxHandler() *PostOutboxHandler {
	return &PostOutboxHandler{}
}

func (h *PostOutboxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
