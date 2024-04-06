package discovery

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bellamariz/go-live-without-downtime/internal/client"
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/rs/zerolog/log"
)

type DiscoveryService struct {
	path   string
	period time.Duration
}

func NewService(cfg *config.Config) *DiscoveryService {
	return &DiscoveryService{
		path:   cfg.OutputStreamPath,
		period: cfg.MaxAgePlaylist,
	}
}

// FetchActiveSignals returns a list of active signals
// An active signal is a signal that had its manifest recently updated
func (ds *DiscoveryService) FetchActiveSignals() []string {
	activeSignals := make([]string, 0)

	err := filepath.Walk(ds.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Error().Err(err).Msgf("Failed to walk path: %s\n", path)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".m3u8" {
			if isRecentlyUpdated(info, ds.period) {
				activeSignals = append(activeSignals, getSignalName(path, ds.path))
			}
		}

		return nil
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to walk stream files")
	}

	return activeSignals
}

// FetchActiveSignals returns a list of active packagers
// Packagers are active when commands 'server1' or 'server2' - our mock local packagers - are running
func (ds *DiscoveryService) FetchActivePackagers(cfg *config.Config) []string {
	activePackagers := make([]string, 0)

	packagerPorts := []string{cfg.ServerOnePort, cfg.ServerTwoPort}

	for _, port := range packagerPorts {
		packagerEndpoint := cfg.LocalHost + ":" + port

		if client.Healthcheck(packagerEndpoint) {
			activePackagers = append(activePackagers, packagerEndpoint)
		}
	}

	return activePackagers
}

// ResetSignals reset the signals cache in the reporter server
// to prevent invalid values
func (ds *DiscoveryService) ResetSignals(cfg *config.Config) error {
	client := &http.Client{Timeout: 2 * time.Second}

	endpoint := fmt.Sprintf("%s:%s/ingests", cfg.LocalHost, cfg.ReporterPort)
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Err(err).Msg("Error reset cache in reporter server")
		return errors.New("error reset cache in reporter server")
	}
	return nil
}

func isRecentlyUpdated(fi os.FileInfo, period time.Duration) bool {
	return time.Since(fi.ModTime()) <= period
}

func getSignalName(path, prefix string) string {
	assetsPathDir := fmt.Sprintf("%s/", prefix)
	manifestDir := strings.TrimPrefix(path, assetsPathDir)
	signal := strings.TrimSuffix(manifestDir, "/playlist.m3u8")

	return signal
}
