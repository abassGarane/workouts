package main

import (
	"net/http"

	"github.com/abassGarane/muscles/domain"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (h *handler) updateWorkout(c echo.Context) error {
	var workout domain.Workout

	// err := json.NewDecoder(c.Request().Body).Decode(&workout)
	if err := echo.New().JSONSerializer.Deserialize(c, &workout); err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, "Unable to decode workout").Error())
	}
	id := c.Param("id")
	returnedWorkout, err := h.service.UpdateWorkout(id, &workout)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, "Unable to decode workout").Error())
	}
	return c.JSON(http.StatusOK, returnedWorkout)
}
