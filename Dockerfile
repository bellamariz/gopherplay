FROM golang:1.22 AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o go-live main.go

FROM linuxserver/ffmpeg
WORKDIR /app
COPY --from=builder /src/ /app/

ENTRYPOINT [ "./go-live" ]