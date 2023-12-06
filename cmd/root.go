package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type PublicIp struct {
	IP string `json:"ip"`
}

var rootCmd = &cobra.Command{
	Use:   "Vpn Checker",
	Short: "Periodically checks if you are still behind a vpn",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
