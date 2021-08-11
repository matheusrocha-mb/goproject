package cmd

import (
	"fmt"
	"os"

	"github.com/matheusrocha-mb/goproject/internal/configuration"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use: "kyc-fees",
	Short: "API to support the management of crypto wallets",
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("Initializing configuration...")

	cobra.OnInitialize(configuration.InitConfig)
	
}