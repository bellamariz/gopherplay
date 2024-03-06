#!/bin/bash

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
output_dir="output"

# Gerando a playlist HLS
ffmpeg -i "$input_video" \
-c:v libx264 -profile:v baseline -level 3.0 -vf scale=-1:480 -maxrate "$video_bitrate" -bufsize "$video_bitrate" -b:a "$audio_bitrate" -c:a aac -strict -2 -f hls \
-segment_time "$segment_duration" -segment_list_size 0 -segment_list "$output_dir/$playlist_name" \
"$output_dir/%05d.ts"

# Movendo a playlist HLS para a pasta de saída
mv "$playlist_name" "$output_dir"

# Exibindo informações sobre a playlist HLS
ffprobe -v quiet -select_streams v:0 -show_entries stream=width,height,bitrate "$output_dir/%05d.ts"

echo "Playlist HLS gerada com sucesso!"