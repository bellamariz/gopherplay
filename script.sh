#!/bin/sh
mkdir output
# Caminho do vídeo MP4 de entrada
input_video="./assets/BigBuckBunny.mp4"

# Nome da playlist HLS
playlist_name="playlist.m3u8"

# Segmentação do vídeo (duração dos segmentos em segundos)
segment_duration=10

# Nível de taxa de bits (bitrate) para o vídeo
video_bitrate="1000k"

# Nível de taxa de bits (bitrate) para o áudio
audio_bitrate="128k"

# Pasta de saída para a playlist HLS
output_dir="./output"

# Gerando a playlist HLS
ffmpeg -stream_loop -1 \
-i "$input_video" \
-c copy \
-f hls \
-strftime 1 -hls_segment_filename "output/segment_%s.ts" \
./output/playlist.m3u8
# -map 1:a -c:v libx264 -x264-params keyint=150:min-keyint=150:scenecut=-1 

echo "Playlist HLS gerada com sucesso!"