package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bellamariz/go-live-without-downtime/internal/worker"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	if err := worker.Execute(ctx); err != nil {
		fmt.Println(err.Error())
		log.Default().Println("Failed generate playlist")
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
