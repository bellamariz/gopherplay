/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/bellamariz/go-live-without-downtime/internal/worker"
	"github.com/spf13/cobra"
)

func RunWorker() *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "Run worker to generate mosaics",
		Run: func(*cobra.Command, []string) {
			ctx := context.Background()

			if err := worker.Execute(ctx); err != nil {
				fmt.Println(err.Error())
				log.Error().Err(err).Msg("Failed to generate playlist")
			}
		},
	}
}
