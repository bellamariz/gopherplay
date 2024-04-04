package api

import (
	"fmt"
	"net/http"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/reporter/ingests"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type API struct {
	Echo *echo.Echo
	Port string
}

func (api *API) healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}

func NewServer(cfg *config.Config) *API {
	return &API{
		Echo: echo.New(),
		Port: cfg.ReporterPort,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/healthcheck", api.healthcheck)
	api.Echo.GET("/ingests", api.getIngests)
	api.Echo.POST("/ingests", api.updateIngests)
}

func (api *API) Start() error {
	return api.Echo.Start(":" + api.Port)
}

func (api *API) getIngests(c echo.Context) error {
	ingests, err := ingests.GetIngests()

	if err != nil {
		log.Error().Err(err).Msg("Error while getting ingest info")

		errMsg := map[string]string{
			"error": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, ingests)
}

func (api *API) updateIngests(c echo.Context) error {
	var ingestSource ingests.Ingest

	if err := c.Bind(&ingestSource); err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, errorMsg)
	}

	if err := c.Validate(ingestSource); err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, errorMsg)
	}

	err := ingests.UpdateIngest(ingestSource)
	if err != nil {
		msg := fmt.Sprintf("Error updating available ingest: %s", err.Error())
		errorMsg := map[string]string{
			"error": msg,
		}

		return c.JSON(http.StatusInternalServerError, errorMsg)
	}

	return c.NoContent(http.StatusOK)
}
