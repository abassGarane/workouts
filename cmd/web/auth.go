package main

import (
	"net/http"

	"github.com/abassGarane/muscles/domain"
	"github.com/abassGarane/muscles/domain/models"
	"github.com/abassGarane/muscles/pkg/passwords"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandler struct {
	s domain.Service
}

func NewAuthHandler(s domain.Service) *AuthHandler {
	return &AuthHandler{s}
}

func (a *AuthHandler) login(c echo.Context) error {
	var user = &models.User{}
	user.ID = primitive.NewObjectID()
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	hashed, _ := passwords.CreateHashedPassword(password)
	user.Name = username
	user.HashedPassword = hashed
	user.Admin = false
	user.Email = email
	return c.JSON(http.StatusOK, echo.Map{"user": user})
}
func (a *AuthHandler) signup(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "signup"})
}
