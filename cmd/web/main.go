package main

import (
	// "log"
	// "net/http"

	"github.com/abassGarane/muscles/domain"
	"github.com/abassGarane/muscles/ui"
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
	service = domain.NewService(repo)
	// if err := http.ListenAndServe(PORT.(string), router); err != nil {
	// 	log.Fatal(err)
	// }
	e := echo.New()
	corsConfig := &middleware.CORSConfig{
		AllowOrigins:     []string{"http://*", "https://*", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "UPDATE", "DELETE", "PATCH"},
		AllowCredentials: false,
	}
	loggerConfig := &middleware.LoggerConfig{
		Format: `"host":"${host}","method":"${method}","uri":"${uri}"`,
	}
	e.Use(middleware.CORSWithConfig(*corsConfig))
	e.Use(middleware.LoggerWithConfig(*loggerConfig))
	apiRouter := e.Group("/api")
	initRouter(service, apiRouter)
	ui.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(PORT.(string)))
}
