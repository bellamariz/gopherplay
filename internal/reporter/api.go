package reporter

import (
	"net/http"
	"sync"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/sources"
	"github.com/labstack/echo/v4"
)

type (
	API struct {
		Echo  *echo.Echo
		Port  string
		Cache *sync.Map
	}
)

func NewServer(cfg *config.Config) *API {
	return &API{
		Echo:  echo.New(),
		Port:  cfg.ReporterPort,
		Cache: &sync.Map{},
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/healthcheck", api.healthcheck)
	api.Echo.GET("/ingests", api.getIngests)
	api.Echo.GET("/ingests/:name", api.getSignalIngest)
	api.Echo.POST("/ingests", api.updateIngests)
}

func (api *API) Start() error {
	return api.Echo.Start(":" + api.Port)
}

func (api *API) healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}

func (api *API) getIngests(c echo.Context) error {
	ingests := []sources.Ingest{}

	api.Cache.Range(func(key, value any) bool {
		ingest := value.(sources.Ingest)

		if ingest.IsActive() {
			ingests = append(ingests, value.(sources.Ingest))
		}

		return true
	})

	if len(ingests) <= 0 {
		errMsg := map[string]string{
			"error": "No active ingest info available",
		}

		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, ingests)
}

func (api *API) updateIngests(c echo.Context) error {
	var ingestSource sources.Ingest

	if err := c.Bind(&ingestSource); err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, errorMsg)
	}

	api.Cache.Store(ingestSource.Signal, ingestSource)

	return c.NoContent(http.StatusOK)
}

func (api *API) getSignalIngest(c echo.Context) error {
	signalName := c.Param("name")

	ingest, found := api.Cache.Load(signalName)
	if !found {
		errorMsg := map[string]string{
			"error": "Active ingests servers not found",
		}

		return c.JSON(http.StatusNotFound, errorMsg)
	}

	return c.JSON(http.StatusOK, ingest)
}
