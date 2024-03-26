#!/bin/sh
mkdir -p output

# Gerando a playlist HLS
ffmpeg \
  -loglevel info \
  -stream_loop -1 -i "assets/globoplay-ad.mp4" \
  -c:v libx264 -profile:v high \
  -c:a copy \
  -f hls \
  -hls_time 5 \
  -hls_list_size 5 \
  -hls_flags delete_segments \
  -strftime 1 -hls_segment_filename "output/seg_%s.ts" \
  "output/playlist.m3u8"

echo "Playlist HLS gerada com sucesso!"