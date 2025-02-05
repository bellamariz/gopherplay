/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/worker"
	"github.com/spf13/cobra"
)

func RunWorker(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "Run worker to generate ffmpeg live stream",
		Run: func(*cobra.Command, []string) {
			ctx := context.Background()

			log.Info().Msg("Running worker for ffmpeg...")

			if err := worker.Execute(ctx, cfg); err != nil {
				log.Error().Err(err).Msg("Failed to generate playlist")
			}
		},
	}
}
