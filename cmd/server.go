package cmd

import (
	"fmt"
	"kobutor/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// ServerCmd ...
	ServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Starts the kobutor server",
		Run:   startServer,
	}
)

func init() {
	ServerCmd.Flags().IntP("port", "p", viper.GetInt("app.port"), "The port to run the kobutor server on")
	viper.BindPFlag("port", ServerCmd.Flags().Lookup("port"))

}

func startServer(cmd *cobra.Command, args []string) {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatal("Failed to start kobutor server")
		}
	}()

	fmt.Println("Starting kobutor server...")

	log.Printf("Kobutor server running at localhost:%d...\n", viper.GetInt("port"))

	<-stop

	log.Println("Kobutor server gracefully shutdown")
}
