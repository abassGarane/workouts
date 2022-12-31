package main

import (
	"github.com/abassGarane/muscles/domain"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func initRouter(s domain.Service) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	// personal middleware

	router.Use(loggingMiddleware)

	//create a handler
	h := NewHandler(s)
	// Real system routes
	router.Get("/health", h.health)
	router.Get("/{id}", h.getWorkout)
	router.Get("/", h.getWorkouts)

	router.Post("/", h.createWorkout)

	router.Delete("/{id}", h.deleteWorkout)

	router.Patch("/{id}", h.updateWorkout)

	return router
}
