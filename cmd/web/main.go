package main

import (
	"github.com/abassGarane/muscles/domain"
	j "github.com/abassGarane/muscles/pkg/jwt"
	"github.com/abassGarane/muscles/ui"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	service domain.Service
)

func main() {
	env := initEnv()
	PORT := env["PORT"]

	repo := initDB()
	service = domain.Newservice(repo)

	e := echo.New()

	corsConfig := &middleware.CORSConfig{
		AllowOrigins:     []string{"http://*", "https://*", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "UPDATE", "DELETE", "PATCH"},
		AllowCredentials: false,
	}
	loggerConfig := &middleware.LoggerConfig{
		Format: `"host":"${host}","method":"${method}","uri":"${uri}"`,
	}

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(j.Claim)
		},
		SigningKey: []byte("jhkafS8AsrFVSAZXFAAG"),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(echo.ErrUnauthorized.Code, echo.Map{
				"message": "User not authorized :: " + err.Error(),
			})
		},
	}

	e.Use(middleware.CORSWithConfig(*corsConfig))
	e.Use(middleware.LoggerWithConfig(*loggerConfig))

	apiRouter := e.Group("/api")
	apiRouter.Use(echojwt.WithConfig(config))
	apiRouter.Use(isAuthenticated)
	authRouter := e.Group("/auth")
	initAuthRouter(service, authRouter)
	initRouter(service, apiRouter)
	ui.RegisterHandlers(e)

	go e.Logger.Fatal(e.Start(PORT.(string)))
}
