package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

var Calls int64

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("Content-Type", "application/json")
		Calls += 1
		fmt.Println("You have access")
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
