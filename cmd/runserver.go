package cmd

import (
	"fmt"

	"github.com/matheusrocha-mb/goproject/internal/database"
	"github.com/spf13/cobra"
)

func runServer(cmd *cobra.Command, args []string) {
	fmt.Println("Initializing database...")
	database.ConnectDB()
	
}

var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Run the gRPC transaction matching server.",
	Run:   runServer,
}

func init() {
	
}
