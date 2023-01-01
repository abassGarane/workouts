package main

import (
	"github.com/abassGarane/muscles/domain"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func initRouter(s domain.Service) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	//cors config
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// personal middleware

	router.Use(loggingMiddleware)

	//create a handler
	h := NewHandler(s)
	// Real system routes
	router.Get("/health", h.health)
	router.Get("/", h.getWorkouts)
	router.Get("/{id}", h.getWorkout)

	router.Post("/", h.createWorkout)

	router.Delete("/{id}", h.deleteWorkout)

	router.Patch("/{id}", h.updateWorkout)

	return router
}
