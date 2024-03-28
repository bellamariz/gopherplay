# Live Streaming Without Downtime (and with Go)

The purpose of this project is to set up a zero downtime strategy for a Live Stream event.

## Running the Live Stream

First, we seek to emulate a live streaming platform flow, as close to real as possible. For this, we use FFMPEG + Golang to produce an HLS playlist and publish it in two separate HTTP servers. This is done by running three different containers:

- `worker` --> Builds and runs the FFMPEG command to generate a live stream output, whose manifest and segments are stored in the folder given by the `OUTPUT_STREAM_PATH` env variable.
- `server1` and `server2` --> Each execute a different HTTP server with equal configurations but different ports (`SERVER_ONE_PORT` and `SERVER_TWO_PORT`).

These services can be run by calling, in different terminals, the following commands:

```sh
go run main.go worker
go run main.go server1
go run main.go server2

```

The output live stream will be an HLS playlist, which can be played in Safari or VLC using the following URLs:

```sh
localhost:8080/playlist.m3u8 ## using server1
localhost:9090/playlist.m3u8 ## using server2
```

The input MP4 video used by the `worker` service to generate the FFMPEG is stored in the folder given by the env variable `INPUT_STREAM_PATH`.

## Zero Downtime Strategy

TBD