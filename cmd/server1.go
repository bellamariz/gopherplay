package cmd

import (
	"github.com/bellamariz/go-live-without-downtime/internal/server"
	"github.com/spf13/cobra"
)

func RunServerOne() *cobra.Command {
	return &cobra.Command{
		Use:   "server1",
		Short: "Run HTTP server",
		Run: func(*cobra.Command, []string) {
			server.Run("8080")
		},
	}
}
