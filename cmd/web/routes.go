package main

import (
	"github.com/abassGarane/muscles/domain"
	// "github.com/go-chi/chi/middleware"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/cors"
	"github.com/labstack/echo/v4"
)

func initRouter(s domain.Service, apiRouter *echo.Group) {
	h := NewHandler(s)
	apiRouter.GET("/health", h.health)
	apiRouter.GET("/", h.getWorkouts)
	apiRouter.GET("/:id", h.getWorkout)

	apiRouter.POST("/", h.createWorkout)
	//
	apiRouter.DELETE("/:id", h.deleteWorkout)

	apiRouter.PATCH("/:id", h.updateWorkout)

}

// func initRouter(s domain.Service) chi.Router {
// 	router := chi.NewRouter()
// 	router.Use(middleware.Logger)
// 	router.Use(middleware.Recoverer)
// 	router.Use(middleware.RequestID)
// 	router.Use(middleware.Heartbeat("/ping"))
// 	router.Use(middleware.AllowContentType("application/json"))
// 	router.Use(middleware.SetHeader("Content-Type", "application/json"))
//
// 	//cors config
// 	router.Use(cors.Handler(cors.Options{
// 		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
// 		AllowedOrigins: []string{"https://*", "http://*", "http://localhost:3000"},
// 		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
// 		ExposedHeaders:   []string{"Link"},
// 		AllowCredentials: false,
// 		MaxAge:           300, // Maximum value not ignored by any of major browsers
// 	}))
//
// 	// personal middleware
//
// 	router.Use(loggingMiddleware)
//
// 	//create a handler
// 	h := NewHandler(s)
// 	// Real system routes
//
// 	apiRouter := chi.NewRouter()
//
// 	apiRouter.Get("/health", h.health)
// 	apiRouter.Get("/", h.getWorkouts)
// 	apiRouter.Get("/{id}", h.getWorkout)
//
// 	apiRouter.Post("/", h.createWorkout)
//
// 	apiRouter.Delete("/{id}", h.deleteWorkout)
//
// 	apiRouter.Patch("/{id}", h.updateWorkout)
//
// 	router.Mount("/api", apiRouter)
// 	return router
// }
