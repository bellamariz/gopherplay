package worker

import "os"

// Build commands to run FFMPEG cli
func BuildCommand() []string {
	orderedArgs := []string{"-loglevel", "debug"}
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
		"-map", "0:v",
		"-map", "0:a",
	}
	return args
}

func buildHLSArguments() []string {
	args := []string{"-f", "hls", "-strftime", "1", "-hls_time", "5",
		"-hls_list_size", "10",
		"-hls_segment_filename", `output/segment_%s.ts`, "output/playlist.m3u8"}
	return args
}
