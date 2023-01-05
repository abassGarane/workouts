package main

import (
	"github.com/abassGarane/muscles/domain"
	"github.com/labstack/echo/v4"
)

func initRouter(s domain.Service, apiRouter *echo.Group) {
	h := NewHandler(s)
	//custom middleware
	apiRouter.Use(loggingMiddleware)

	apiRouter.GET("/health", h.health)
	apiRouter.GET("", h.getWorkouts)
	apiRouter.GET("/:id", h.getWorkout)

	apiRouter.POST("/", h.createWorkout)

	apiRouter.DELETE("/:id", h.deleteWorkout)

	apiRouter.PATCH("/:id", h.updateWorkout)
}
