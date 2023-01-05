package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (h *handler) deleteWorkout(c echo.Context) error {
	id := c.Param("id")
	err = h.service.DeleteWorkout(id)
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, errors.Wrap(err, "Unable to delete workout").Error())
	}
	return c.NoContent(http.StatusOK)
}
