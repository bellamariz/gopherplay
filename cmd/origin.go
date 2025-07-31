package cmd

import (
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/origin"
	"github.com/spf13/cobra"
)

func RunOrigin(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "origin",
		Short: "Run origin server that awnswer the active signal server",
		Run: func(*cobra.Command, []string) {
			originAPI := origin.NewServer(cfg)
			originAPI.ConfigureRoutes()
			err := originAPI.Start()
			if err != nil {
				panic("failed to start origin service: " + err.Error())
			}
		},
	}
}
