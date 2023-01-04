package ui

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed all:build
	build embed.FS
	//go:embed build/index.html
	index embed.FS

	buildDirFs     = echo.MustSubFS(build, "build")
	buildIndexHtml = echo.MustSubFS(index, "build")
)

func RegisterHandlers(e *echo.Echo) {
	e.FileFS("/", "index.html", buildIndexHtml)
	e.StaticFS("/", buildDirFs)
}
