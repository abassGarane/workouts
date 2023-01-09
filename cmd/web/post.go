package main

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/abassGarane/muscles/domain"
	j "github.com/abassGarane/muscles/pkg/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (h *handler) createWorkout(c echo.Context) error {
	var workout domain.Workout
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*j.Claim)

	err := c.Bind(&workout)
	workout.UserEmail = claims.Email
	c.Logger().Debug(workout)
	if err := domain.Validate(&workout); err != nil {
		return c.JSON(echo.ErrBadRequest.Code, err.Error())
	}
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, errors.Wrap(err, fmt.Errorf("could not decode the request body %v", err).Error()).Error())
	}

	work, err := h.service.CreateWorkout(&workout)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, fmt.Errorf("Could not encode the request body %v", err).Error()).Error())
	}

	return c.JSON(http.StatusOK, work)
}
