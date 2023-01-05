package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (h *handler) health(c echo.Context) error {
	// w.Header().Set("Content-Type", "application/json") // for tests
	return c.JSON(http.StatusOK, echo.Map{
		"status": "running...",
	})
}

func (h *handler) getWorkouts(c echo.Context) error {
	workouts, err := h.service.GetWorkouts()
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, "Unable to retrieve workouts").Error())
	}
	return c.JSON(http.StatusOK, workouts)
}

func (h *handler) getWorkout(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(c.ParamValues())
	wkout, err := h.service.GetWorkout(id)
	fmt.Println("Retrieved workout :: GetWorkout", wkout)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, "Unable to retrieve workout").Error())
	}
	return c.JSON(http.StatusOK, wkout)
}
