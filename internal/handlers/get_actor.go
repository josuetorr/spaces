package handlers

import (
	"net/http"
)

type GetActorHandler struct{}

func NewGetActorHandler() *GetActorHandler {
	return &GetActorHandler{}
}

func (h *GetActorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
