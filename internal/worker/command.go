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
	args := []string{"-c", "copy"}
	return args
}

func buildHLSArguments() []string {
	args := []string{"-f", "hls", "-strftime", "1",
		"-hls_start_number_source", "epoch", "-hls_segment_filename", "output/segment_%s.ts", "output/playlist.m3u8"}
	return args
}
