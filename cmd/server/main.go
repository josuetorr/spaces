package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
	"gitlab.com/josuetorr/spaces/internal/data"
	"gitlab.com/josuetorr/spaces/internal/handlers"
	"gitlab.com/josuetorr/spaces/internal/services"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdin, nil))
	db, err := data.Init()
	if err != nil {
		panic(err)
	}

	actorRepo := data.NewActorRepo(db)
	actorService := services.NewActorService(actorRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/.well-known/webfinger", handlers.NewWebFingerHandler().ServeHTTP)

	r.Post("/users/", handlers.NewPostActorHandler(log, actorService).ServeHTTP)
	r.Get("/users/{username}", handlers.NewGetActorHandler().ServeHTTP)

	r.Get("/users/{username}/inbox", handlers.NewGetInboxHandler().ServeHTTP)
	r.Post("/users/{username}/inbox", handlers.NewPostInboxHandler(log).ServeHTTP)

	r.Get("/users/{username}/outbox", handlers.NewGetOutboxHandler().ServeHTTP)
	r.Post("/users/{username}/outbox", handlers.NewPostOutboxHandler().ServeHTTP)

	port := "3000"

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			_ = fmt.Errorf("Server error: %s", err)
		}
	}()

	fmt.Println("Server started on :", port)

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}

	fmt.Println("Server exited.")
}
