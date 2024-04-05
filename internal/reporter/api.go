package reporter

import (
	"net/http"
	"sync"

	"github.com/go-playground/validator"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type (
	API struct {
		Echo  *echo.Echo
		Port  string
		Cache *sync.Map
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func NewServer(cfg *config.Config) *API {
	server := echo.New()
	server.Validator = &CustomValidator{validator: validator.New()}

	return &API{
		Echo:  echo.New(),
		Port:  cfg.ReporterPort,
		Cache: &sync.Map{},
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

func (api *API) healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}

func (api *API) getIngests(c echo.Context) error {
	ingests := []Ingest{}

	api.Cache.Range(func(key, value any) bool {
		ingests = append(ingests, value.(Ingest))
		return true
	})

	if len(ingests) <= 0 {
		log.Error().Msg("Error while getting ingest info")

		errMsg := map[string]string{
			"error": "No ingest info available in cache",
		}

		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusOK, ingests)
}

func (api *API) updateIngests(c echo.Context) error {
	var ingestSource Ingest

	if err := c.Bind(&ingestSource); err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, errorMsg)
	}

	api.Cache.Store(ingestSource.Signal, ingestSource)

	return c.NoContent(http.StatusOK)
}
