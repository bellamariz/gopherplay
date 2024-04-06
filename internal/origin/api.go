package origin

import (
	"net/http"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/labstack/echo/v4"
)

type (
	API struct {
		Echo             *echo.Echo
		Port             string
		ReporterEndpoint string
	}
)

func NewServer(cfg *config.Config) *API {
	return &API{
		Echo:             echo.New(),
		Port:             cfg.OriginPort,
		ReporterEndpoint: cfg.LocalHost + ":" + cfg.ReporterPort,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/healthcheck", api.healthcheck)
	api.Echo.GET("/live/:name", api.getSignalServer)
	api.Echo.GET("/signals", api.getSignals)
}

func (api *API) Start() error {
	return api.Echo.Start(":" + api.Port)
}

func (api *API) healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}

func (api *API) getSignals(c echo.Context) error {
	signals, err := listSignals(api.ReporterEndpoint)
	if err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMsg)
	}

	return c.JSON(http.StatusOK, signals)
}

func (api *API) getSignalServer(c echo.Context) error {
	name := c.Param("name")

	signalInfo, err := getSignalPackagers(api.ReporterEndpoint, name)
	if err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMsg)
	}

	if len(signalInfo.Packagers) == 0 {
		if err != nil {
			errorMsg := map[string]string{
				"error": "The signal does not have any active packager as ingest",
			}
			return c.JSON(http.StatusBadRequest, errorMsg)
		}
	}

	activeSignalPath := formatPath(signalInfo.Packagers, signalInfo.Signal)
	return c.JSON(http.StatusOK, activeSignalPath)
}
