package cmd

import (
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/spf13/cobra"
)

func RunReporter(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "reporter",
		Short: "Run reporter service to expose active ingests",
		Run: func(*cobra.Command, []string) {
		},
	}
}
