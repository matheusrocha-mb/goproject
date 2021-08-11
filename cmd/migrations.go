package cmd

import (
	"fmt"

	"github.com/matheusrocha-mb/goproject/internal/database"
	"github.com/spf13/cobra"
)

func runMigration(cmd *cobra.Command, args []string) {
	if err := database.ConnectDB(); err != nil {
		fmt.Printf("Can't connect database : %v", err)
		return
	}
	database.Migrate()
}

var migrateCMD = &cobra.Command{
	Use: "migrate",
	Short: "Perform DB migration",
	Run: runMigration,
}

func init() {

}