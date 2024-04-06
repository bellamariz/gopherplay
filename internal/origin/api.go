package origin

import (
	"net/http"
	"sync"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	API struct {
		Echo           *echo.Echo
		Port           string
		ReportEndPoint string
		Cache          *sync.Map
	}

	ServerParams struct {
		Config           *config.Config
		ReporterEndpoint string
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

func NewServer(params ServerParams) *API {
	server := echo.New()
	server.Validator = &CustomValidator{validator: validator.New()}

	return &API{
		Echo:           echo.New(),
		Port:           params.Config.OriginPort,
		ReportEndPoint: params.ReporterEndpoint,
		Cache:          &sync.Map{},
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
	signals, err := listSignals(api.ReportEndPoint)
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

	signalInfo, err := getSignalPackagers(api.ReportEndPoint, name)
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
