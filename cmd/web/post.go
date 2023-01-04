package main

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/abassGarane/muscles/domain"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (h *handler) createWorkout(c echo.Context) error {
	var workout domain.Workout

	// err := json.NewDecoder(c.Request().Body).Decode(&workout)
	err := echo.New().JSONSerializer.Deserialize(c, &workout)
	c.Logger().Debug(workout)

	if err != nil {
		return c.String(echo.ErrBadRequest.Code, errors.Wrap(err, fmt.Errorf("could not decode the request body %v", err).Error()).Error())
	}

	work, err := h.service.CreateWorkout(&workout)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, fmt.Errorf("Could not encode the request body %v", err).Error()).Error())
	}

	return c.JSON(http.StatusOK, work)
}
