package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abassGarane/muscles/domain"
	"github.com/abassGarane/muscles/domain/models"
	j "github.com/abassGarane/muscles/pkg/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type AuthHandler struct {
	s domain.Service
}

func NewAuthHandler(s domain.Service) *AuthHandler {
	return &AuthHandler{s}
}

func (a *AuthHandler) login(c echo.Context) error {
	l := &models.UserLoginRequest{}
	if err := c.Bind(l); err != nil {
		if l.Validate() != nil {
			return c.JSON(http.StatusOK, echo.Map{"message": errors.Wrap(err, "all fields should not be empty")})
		}
		return c.JSON(http.StatusOK, echo.Map{"message": err})
	}
	fmt.Println("Login request::", l)
	user, err := a.s.GetUserByEmail(l.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User does not exist"})
	}
	//check password
	err = j.Verify(user.HashedPassword, l.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Wrong password"})
	}
	claims := j.Claim{
		Email:    user.Email,
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}
	token := j.CreateSignature(claims, "the world is a beautiful place")
	return c.JSON(200, echo.Map{
		"token": token,
	})
}
func (a *AuthHandler) signup(c echo.Context) error {
	user := &models.UserRequest{}
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": fmt.Sprintf("Validation error %v", err)})
	}
	fmt.Println(user)
	err = user.Validate()
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": fmt.Sprintf("Could not create user %v", err)})
	}
	//see if user exists
	_, err = a.s.GetUserByEmail(user.Email)
	if err == nil {
		return c.JSON(401, echo.Map{"message": "user already exists"})
	}
	createdUser := &models.User{}
	createdUser, err = a.s.CreateUser(user)
	fmt.Println(createdUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": fmt.Sprintf("Could not create user %v", err)})
	}
	fmt.Println(createdUser.Email)
	claims := j.Claim{
		Username: createdUser.Name,
		Email:    createdUser.Email,
		Admin:    false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}
	fmt.Println("Claims")
	fmt.Println(claims)
	token := j.CreateSignature(claims, "the world is a beautiful place")
	fmt.Println(token)
	return c.JSON(200, echo.Map{
		"token": token,
		"user":  createdUser,
	})
	// return c.JSON(http.StatusOK, echo.Map{"user": createdUser})
}
