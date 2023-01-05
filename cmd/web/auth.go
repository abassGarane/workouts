package main

import (
	"net/http"
	"time"

	"github.com/abassGarane/muscles/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const secret = "aiahihsojaojdaykdxdAYFD8IQU23RWEDFVASUJHXAYUAFT7"

type AuthHandler struct {
	service domain.Service
}

func NewAuthHandler(s domain.Service) *AuthHandler {
	return &AuthHandler{s}

}

func (a *AuthHandler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return echo.ErrUnauthorized
	}
	claims := &Claim{
		Name:  username,
		ID:    uuid.NewString(),
		Admin: false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
