package main

import (
	"github.com/abassGarane/muscles/domain"
	"github.com/labstack/echo/v4"
)

func initAuthRouter(s domain.Service, authRouter *echo.Group) {

	h := NewAuthHandler(s)

	authRouter.POST("/login", h.login)
	authRouter.POST("/signup", h.signup)
}
