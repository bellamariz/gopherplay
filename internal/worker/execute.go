package worker

import (
	"context"
	"os"
	"os/exec"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/rs/zerolog/log"
)

// Execute FFMPEG command, create HLS playlist from
// mp4 video
func Execute(ctx context.Context, cfg *config.Config) error {
	outputPath := cfg.OutputStreamPath + "/" + cfg.LiveSignalName
	if err := CreateOutputDir(outputPath); err != nil {
		log.Error().Err(err).Msg("Failed to create output directory for generated playlist")
		return err
	}

	args := BuildCommand(cfg)
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func CreateOutputDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
