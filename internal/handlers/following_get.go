package handlers

import "net/http"

type GetFollowingHandler struct{}

func NewGetFollowingHandler() *GetFollowingHandler {
	return &GetFollowingHandler{}
}

func (h *GetFollowingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
