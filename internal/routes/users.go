package routes

import (
	"log/slog"

	"github.com/go-chi/chi"
	"gitlab.com/josuetorr/spaces/internal/handlers"
)

func NewUserRoutes(actorService handlers.ActorService, log *slog.Logger) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handlers.NewPostActorHandler(log, actorService).ServeHTTP)

	return r
}
