package handlers

import "net/http"

type GetFollowersHandler struct{}

func NewGetFollowersHandler() *GetFollowersHandler {
	return &GetFollowersHandler{}
}

func (h *GetFollowersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
