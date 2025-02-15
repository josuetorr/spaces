package routes

import (
	"log/slog"

	"github.com/go-chi/chi"
	"gitlab.com/josuetorr/spaces/internal/handlers"
)

func NewFederationRoutes(actorService handlers.ActorService, log *slog.Logger) chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.NewGetActorHandler(actorService).ServeHTTP)

	r.Get("/inbox", handlers.NewGetInboxHandler().ServeHTTP)
	r.Post("/inbox", handlers.NewPostInboxHandler(log).ServeHTTP)

	r.Get("/outbox", handlers.NewGetOutboxHandler().ServeHTTP)
	r.Post("/outbox", handlers.NewPostOutboxHandler().ServeHTTP)

	r.Get("/following", handlers.NewGetFollowingHandler().ServeHTTP)
	r.Get("/followers", handlers.NewGetFollowersHandler().ServeHTTP)

	return r
}
