package cmd

import (
	"kobutor/config"
	"log"

	"github.com/spf13/cobra"
)

var (
	// RootCmd ...
	RootCmd = &cobra.Command{
		Use:   "kobutor",
		Short: "kobutor is an http service for sending emails",
	}
)

func init() {
	config.Init()
	RootCmd.AddCommand(ServerCmd)
}

// Execute executes the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
