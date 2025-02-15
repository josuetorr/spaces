package routes

import (
	"log/slog"

	"github.com/dgraph-io/dgo/v240"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gitlab.com/josuetorr/spaces/internal/data"
	"gitlab.com/josuetorr/spaces/internal/handlers"
	"gitlab.com/josuetorr/spaces/internal/services"
)

func SetupRoutes(dg *dgo.Dgraph, log *slog.Logger) chi.Router {
	actorRepo := data.NewActorRepo(dg, log)
	actorService := services.NewActorService(actorRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/.well-known/webfinger", handlers.NewWebFingerHandler(log, actorService).ServeHTTP)

	r.Mount("/users", NewUserRoutes(actorService, log))
	r.Mount("/users/{username}", NewFederationRoutes(actorService, log))

	return r
}
