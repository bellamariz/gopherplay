package server

import (
	"net/http"
	"path/filepath"

	"github.com/bellamariz/go-live-without-downtime/internal/mimetype"
	"github.com/labstack/echo/v4"
)

func Run(port, outputStreamPath string) {
	mimetype.Configure()

	e := echo.New()

	e.GET("/healthcheck", healthCheck)
	e.GET("/*", serveStatic(outputStreamPath))

	e.Logger.Fatal(e.Start(":" + port))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}

func serveStatic(root string) echo.HandlerFunc {
	return func(c echo.Context) error {
		file := filepath.Join(root, c.Request().URL.Path)

		return c.File(file)
	}
}
