package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number for vpn_checker",
	Long:  `All software has versions. This is vpn_checker`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vpn_checker version 1.0")
	},
	Args: cobra.NoArgs,
}
