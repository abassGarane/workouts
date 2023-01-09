package main

import (
	"context"
	"fmt"

	j "github.com/abassGarane/muscles/pkg/jwt"
	"github.com/golang-jwt/jwt/v4"
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

func isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*j.Claim)
		if claims.Valid() == nil {
			c.Request().WithContext(context.WithValue(c.Request().Context(), "user_email", claims.Email))
			next(c)
		}
		return echo.ErrUnauthorized
	}

}
