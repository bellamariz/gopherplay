package cmd

import (
	"fmt"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/origin"
	"github.com/spf13/cobra"
)

func RunOrigin(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "origin",
		Short: "Run origin server that awnswer the active signal server",
		Run: func(*cobra.Command, []string) {
			fmt.Println("Origin server =)")
			reporterUrl := fmt.Sprintf("%s:%s", cfg.LocalHost, cfg.ReporterPort)
			serverParams := origin.ServerParams{
				Config:           cfg,
				ReporterEndpoint: reporterUrl,
			}

			originServer := origin.NewServer(serverParams)

			originServer.ConfigureRoutes()
			originServer.Start()
		},
	}
}
