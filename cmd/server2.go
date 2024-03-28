package cmd

import (
	"github.com/bellamariz/go-live-without-downtime/internal/server"
	"github.com/spf13/cobra"
)

func RunServerTwo() *cobra.Command {
	return &cobra.Command{
		Use:   "server2",
		Short: "Run HTTP server",
		Run: func(*cobra.Command, []string) {
			server.Run("9090")
		},
	}
}
