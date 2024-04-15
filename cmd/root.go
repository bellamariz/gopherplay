/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewRootCmd(cfg *config.Config) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "go-live",
		Short:         "Run full framework for go live application",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(RunServerOne(cfg), RunServerTwo(cfg), RunWorker(cfg), RunDiscovery(cfg), RunReporter(cfg), RunOrigin(cfg), RunProxy(cfg))

	return rootCmd
}

func Execute() {
	cfg := SetupEnvironment()

	if err := NewRootCmd(cfg).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetupEnvironment() *config.Config {
	// Configure global logging
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime})

	// Read env file (for running locally)
	if err := godotenv.Load(".env"); err != nil {
		log.Info().Err(err).Msg("Could not load .env file")
	}

	// Parse environment configuration variables
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to process environment configurations")
	}

	// Create local folder for storing the ffmpeg mosaic output
	if err := os.MkdirAll(cfg.OutputStreamPath, os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("Failed to create output path directory")
	}

	return cfg
}
