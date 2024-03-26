package worker

import (
	"context"
	"os"
	"os/exec"
)

// Execute FFMPEG command, create HLS playlist from
// mp4 video
func Execute(ctx context.Context) error {
	args := BuildCommand()
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
