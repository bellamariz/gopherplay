#!/bin/sh
mkdir -p output

# Generating HLS playlist
ffmpeg \
  -loglevel info \
  -stream_loop -1 -i "assets/stream.mp4" \
  -c:v libx264 -profile:v high \
  -c:a copy \
  -f hls \
  -hls_time 5 \
  -hls_list_size 5 \
  -hls_flags delete_segments \
  -strftime 1 -hls_segment_filename "output/seg_%s.ts" \
  "output/playlist.m3u8"

echo "Playlist generated successfully!"