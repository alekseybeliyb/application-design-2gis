package main

import (
	"app-booking/internal/handler"
	"app-booking/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	router := chi.NewRouter()

	repo := repository.NewInMemoryRepository()
	repo.Migrate()

	orderHandler := handler.NewOrderHandler(repo)

	router.Use(middleware.Recoverer)

	router.Post("/orders", orderHandler.CreateOrder)
	router.Get("/orders/all", orderHandler.GetOrders)

	slog.Info("Server listening localhost:8080")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
