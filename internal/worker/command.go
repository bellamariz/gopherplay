package worker

import (
	"fmt"
	"os"
)

// Build commands to run FFMPEG cli
func BuildCommand() []string {
	orderedArgs := []string{"-loglevel", "info"}

	orderedArgs = append(orderedArgs, buildVideoInputArguments()...)
	orderedArgs = append(orderedArgs, buildCodecsConfig()...)
	orderedArgs = append(orderedArgs, buildHLSArguments()...)

	return orderedArgs
}

func buildVideoInputArguments() []string {
	args := []string{"-stream_loop", "-1", "-i"}

	args = append(args, os.Getenv("VIDEO_PATH"))

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

func buildHLSArguments() []string {
	segmentPattern := fmt.Sprintf("output/seg_%%s.ts")

	args := []string{
		"-f", "hls",
		"-hls_time", "5",
		"-hls_list_size", "10",
		"-hls_flags", "delete_segments",
		"-strftime", "1",
		"-hls_segment_filename", segmentPattern, "output/playlist.m3u8",
	}

	return args
}
