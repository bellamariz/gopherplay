package cmd

import (
	"github.com/bellamariz/go-live-without-downtime/internal/client"
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/discovery"
	"github.com/bellamariz/go-live-without-downtime/internal/reporter"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func RunDiscovery(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "discovery",
		Short: "Run discovery service to expose active signals to reporter",
		Run: func(*cobra.Command, []string) {
			discoveryService := discovery.NewService(cfg)
			reporterService := reporter.NewService(cfg, discoveryService)

			httpClient := client.New()

			if !httpClient.Healthcheck(reporterService.Endpoint) {
				log.Error().Msg("Reporter service is not running")
				return
			}

			reporterService.Start(cfg)
		},
	}
}
