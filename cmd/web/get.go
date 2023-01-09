package main

import (
	"fmt"
	"net/http"

	j "github.com/abassGarane/muscles/pkg/jwt"
	"github.com/golang-jwt/jwt/v4"
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*j.Claim)
	user_email := claims.Email
	workouts, err := h.service.GetWorkouts(user_email)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, "Unable to retrieve workouts").Error())
	}
	return c.JSON(http.StatusOK, workouts)
}

func (h *handler) getWorkout(c echo.Context) error {
	user_email := c.Request().Context().Value("user").(string)
	fmt.Println(user_email)
	id := c.Param("id")
	fmt.Println(c.ParamValues())
	wkout, err := h.service.GetWorkout(id)
	fmt.Println("Retrieved workout :: GetWorkout", wkout)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, "Unable to retrieve workout").Error())
	}
	return c.JSON(http.StatusOK, wkout)
}
