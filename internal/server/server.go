package server

import (
	"net/http"

	"github.com/bellamariz/go-live-without-downtime/internal/mimetype"
	"github.com/labstack/echo/v4"
)

func Run(port string) {
	mimetype.Configure()

	e := echo.New()

	e.GET("healthcheck", healthCheck)

	e.Logger.Fatal(e.Start(":" + port))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}
