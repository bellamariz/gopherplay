package reporter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bellamariz/go-live-without-downtime/internal/client"
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/discovery"
	"github.com/bellamariz/go-live-without-downtime/internal/sources"
	"github.com/rs/zerolog/log"
)

type ReporterService struct {
	Client          *client.HTTPClient
	PackagerService *discovery.DiscoveryService
	Endpoint        string
}

func NewService(cfg *config.Config, ds *discovery.DiscoveryService) *ReporterService {
	return &ReporterService{
		Client:          client.New(),
		PackagerService: ds,
		Endpoint:        cfg.LocalHost + ":" + cfg.ReporterPort,
	}
}

func (rs *ReporterService) Start(cfg *config.Config) {
	ticker := time.NewTicker(cfg.DiscoveryRunFrequency)

	for range ticker.C {
		rs.SetIngest(cfg)
	}
}

func (rs *ReporterService) SetIngest(cfg *config.Config) {
	activePackagers := rs.PackagerService.FetchActivePackagers(cfg)
	activeSignals := rs.PackagerService.FetchActiveSignals()

	for _, signal := range activeSignals {
		rs.setSignalIngest(signal, activePackagers)
	}
}

func (rs *ReporterService) setSignalIngest(signal string, packagers []string) {
	now := time.Now().Format(time.RFC1123)
	ingestSource := sources.Ingest{Packagers: packagers, Signal: signal, LastReported: now}

	payload, err := json.Marshal(ingestSource)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal signal ingest data")
		return
	}

	endpoint := fmt.Sprintf("%s/ingests", rs.Endpoint)

	err = rs.Client.Post(endpoint, "application/json", payload)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to set ingest for '%s' signal", signal)
		return
	}

	log.Info().Msgf("Registered signal '%s' as ingest source", signal)
}
