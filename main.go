package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bellamariz/go-live-without-downtime/internal/server"
	"github.com/bellamariz/go-live-without-downtime/internal/worker"
)

func main() {
	ctx := context.Background()
	if err := worker.Execute(ctx); err != nil {
		fmt.Println(err.Error())
		log.Default().Println("Failed generate playlist")
	}

	server.Run()
}
