package routes

import (
	"log/slog"

	"github.com/go-chi/chi"
	"gitlab.com/josuetorr/spaces/internal/handlers"
)

func NewFederationRoutes(actorService handlers.ActorService, activityService handlers.ActivityService, log *slog.Logger) chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.NewGetActorHandler(actorService).ServeHTTP)

	r.Get("/inbox", handlers.NewGetInboxHandler().ServeHTTP)
	r.Post("/inbox", handlers.NewPostInboxHandler(log, activityService).ServeHTTP)

	r.Get("/outbox", handlers.NewGetOutboxHandler().ServeHTTP)
	r.Post("/outbox", handlers.NewPostOutboxHandler(log, activityService).ServeHTTP)

	r.Get("/following", handlers.NewGetFollowingHandler(log, actorService).ServeHTTP)
	r.Get("/followers", handlers.NewGetFollowersHandler().ServeHTTP)

	return r
}
