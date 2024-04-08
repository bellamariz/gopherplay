package discovery

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bellamariz/go-live-without-downtime/internal/client"
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/rs/zerolog/log"
)

type DiscoveryService struct {
	httpClient *client.HTTPClient
	path       string
	period     time.Duration
}

func NewService(cfg *config.Config) *DiscoveryService {
	return &DiscoveryService{
		httpClient: client.New(),
		path:       cfg.OutputStreamPath,
		period:     cfg.MaxAgePlaylist,
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

		if ds.httpClient.Healthcheck(packagerEndpoint) {
			activePackagers = append(activePackagers, packagerEndpoint)
		}
	}

	return activePackagers
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
