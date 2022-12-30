package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func initRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.AllowContentType("application/json"))

	// personal middleware

	router.Use(loggingMiddleware)

	// Real system routes
	router.Get("/health", health)
	router.Get("/:id", getWorkout)
	router.Get("/", getWorkouts)

	router.Post("/", createWorkout)

	router.Delete("/:id", deleteWorkout)

	router.Patch("/:id", updateWorkout)

	return router
}
