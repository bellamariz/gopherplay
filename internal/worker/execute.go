package worker

import (
	"context"
	"os"
	"os/exec"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
)

// Execute FFMPEG command, create HLS playlist from
// mp4 video
func Execute(ctx context.Context, cfg *config.Config) error {
	args := BuildCommand(cfg)
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
