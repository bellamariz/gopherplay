#!/bin/sh
mkdir output
# Caminho do vídeo MP4 de entrada
input_video="./assets/BigBuckBunny.mp4"

# Nome da playlist HLS
playlist_name="playlist.m3u8"

# Gerando a playlist HLS
ffmpeg -stream_loop -1 \
-i "./assets/BigBuckBunny.mp4" \
-c:v libx264 -profile:v high -x264-params keyint=150:min-keyint=150:scenecut=-1 \
-map 0:v -map 0:a \
-f hls \
-hls_time 5 \
-hls_segment_type mpegts \
-strftime 1 \
-hls_list_size 10 \
-hls_flags delete_segments+program_date_time+omit_endlist \
-hls_segment_filename "./output/segment_%s.ts" \
"./output/playlist.m3u8"

echo "Playlist HLS gerada com sucesso!"