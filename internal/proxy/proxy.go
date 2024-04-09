package proxy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/bellamariz/go-live-without-downtime/internal/client"
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/sources"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type (
	API struct {
		Echo           *echo.Echo
		Port           string
		OriginEndpoint string
	}
)

func NewProxyServer(cfg *config.Config) *API {
	return &API{
		Echo:           echo.New(),
		Port:           cfg.ProxyPort,
		OriginEndpoint: fmt.Sprintf("%s:%s", cfg.LocalHost, cfg.OriginPort),
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/:name/*", api.proxySignalVideo)
}

func (api *API) Start() error {
	return api.Echo.Start(":" + api.Port)
}

func (api *API) proxySignalVideo(c echo.Context) error {
	fmt.Println("get signal method")
	name := c.Param("name")
	if name == "" {
		errorMsg := map[string]string{
			"error": "should pass a signal name",
		}
		return c.JSON(http.StatusBadRequest, errorMsg)
	}

	originSignal := fmt.Sprintf("%s/live/%s", api.OriginEndpoint, name)

	httpClient := client.New()
	resp, err := httpClient.Get(originSignal)
	if err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errorMsg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errorMsg)
	}

	defer resp.Body.Close()

	var response sources.Source
	if err := json.Unmarshal(body, &response); err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errorMsg)
	}

	log.Info().Msgf("Proxy live to server %s ", response.Server)

	url, _ := url.Parse(response.Server)
	rp := &httputil.ReverseProxy{
		Director: newDirector(url),
	}

	rp.ServeHTTP(c.Response(), c.Request())
	return nil
}

func newDirector(url *url.URL) func(req *http.Request) {
	return func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
	}
}
