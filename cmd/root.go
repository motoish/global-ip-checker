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
	Long: `Vpn Checker will poll a ip api back that will return your current global ip
It will just compared your current global ip with your vpn global ip address. 
If the ip address do not match then you will get a notification that you are currently
not connected your VPN service.`,
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
