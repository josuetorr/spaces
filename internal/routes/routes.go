package routes

import (
	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kivik/kivik/v4"
	"gitlab.com/josuetorr/spaces/internal/data"
	"gitlab.com/josuetorr/spaces/internal/handlers"
	"gitlab.com/josuetorr/spaces/internal/services"
)

func SetupRoutes(db *kivik.DB, log *slog.Logger) chi.Router {
	actorRepo := data.NewActorRepo(log, db)
	activityRepo := data.NewActivityRepo(log, db)

	actorService := services.NewActorService(actorRepo)
	activityService := services.NewActivityService(log, activityRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/.well-known/webfinger", handlers.NewWebFingerHandler(log, actorService).ServeHTTP)

	r.Mount("/users", NewUserRoutes(actorService, log))
	r.Mount("/users/{username}", NewFederationRoutes(actorService, activityService, log))

	return r
}
