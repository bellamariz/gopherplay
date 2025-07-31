package cmd

import (
	"github.com/bellamariz/go-live-without-downtime/internal/config"
	"github.com/bellamariz/go-live-without-downtime/internal/proxy"
	"github.com/spf13/cobra"
)

func RunProxy(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "proxy",
		Short: "Run proxy server to player",
		Run: func(*cobra.Command, []string) {
			proxyAPI := proxy.NewProxyServer(cfg)
			proxyAPI.ConfigureRoutes()
			err := proxyAPI.Start()
			if err != nil {
				panic("failed to start proxy service: " + err.Error())
			}
		},
	}
}
