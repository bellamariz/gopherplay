package reporter

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bellamariz/go-live-without-downtime/internal/client"
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/discovery"
	"github.com/rs/zerolog/log"
)

type Ingest struct {
	Signal       string   `json:"signal"`
	Packagers    []string `json:"packagers"`
	LastReported string   `json:"last_reported"`
}

type ReporterService struct {
	PackagerService *discovery.DiscoveryService
	Endpoint        string
}

func NewService(cfg *config.Config, ds *discovery.DiscoveryService) *ReporterService {
	return &ReporterService{
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

	if len(activePackagers) <= 0 {
		log.Warn().Msg("There are no active packagers")
		return
	}

	activeSignals := rs.PackagerService.FetchActiveSignals()

	if len(activeSignals) <= 0 {
		log.Warn().Msg("There are no active signals")
		return
	}

	for _, signal := range activeSignals {
		rs.setSignalIngest(signal, activePackagers)
	}
}

func (rs *ReporterService) setSignalIngest(signal string, packagers []string) {
	now := time.Now().Format(time.RFC1123)
	ingestSource := Ingest{Packagers: packagers, Signal: signal, LastReported: now}

	payload, err := json.Marshal(ingestSource)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal signal ingest data")
		return
	}

	resp, err := client.Post(rs.Endpoint, "ingests", "application/json", payload)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to set ingest for '%s' signal", signal)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Err(err).Msgf("Failed to set ingest for '%s' signal: got '%s'", signal, resp.Status)
		return
	}

	log.Info().Msgf("Registered signal '%s' as ingest source", signal)
}
