package cmd

import (
	"goqueue/cmd/queue"
	"goqueue/config"
	"log"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "kobutor",
		Short: "kobutor is an http service for sending emails",
	}
)

func init() {
	config.Init()
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(queue.QueueCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
