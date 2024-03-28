/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "application",
		Short:         "Run framework for application",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(RunServerOne(), RunServerTwo(), RunWorker())

	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
