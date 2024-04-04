package worker

import (
	"fmt"

	"github.com/bellamariz/go-live-without-downtime/internal/config"
)

// Build commands to run FFMPEG cli
func BuildCommand(cfg *config.Config) []string {
	orderedArgs := []string{"-loglevel", "info"}

	orderedArgs = append(orderedArgs, buildVideoInputArguments(cfg)...)
	orderedArgs = append(orderedArgs, buildCodecsConfig()...)
	orderedArgs = append(orderedArgs, buildHLSArguments(cfg)...)

	return orderedArgs
}

func buildVideoInputArguments(cfg *config.Config) []string {
	args := []string{
		"-stream_loop", "-1",
		"-i", cfg.InputStreamPath,
	}

	return args
}

func buildCodecsConfig() []string {
	args := []string{
		"-c:v", "libx264",
		"-profile:v", "high",
		"-c:a", "copy",
	}

	return args
}

func buildHLSArguments(cfg *config.Config) []string {
	outputPath := cfg.OutputStreamPath + "/" + cfg.LiveSignalName
	segmentPattern := fmt.Sprintf("%s/seg_%%s.ts", outputPath)
	playlistPath := fmt.Sprintf("%s/playlist.m3u8", outputPath)

	args := []string{
		"-f", "hls",
		"-hls_time", "5",
		"-hls_list_size", "10",
		"-hls_flags", "delete_segments",
		"-strftime", "1",
		"-hls_segment_filename", segmentPattern,
		playlistPath,
	}

	return args
}
