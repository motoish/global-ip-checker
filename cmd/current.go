package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var getCurrentGlobalIpCmd = &cobra.Command{
	Use:   "current",
	Short: "Get your current global ip",
	Run:   getCurrentGlobalIp,
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(getCurrentGlobalIpCmd)
}

func getCurrentGlobalIp(cmd *cobra.Command, args []string) {
	response, error := http.Get(ipify)
	if error != nil {
		fmt.Println("Error Checking your global ip")
		return
	}

	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println("Error reading initial response body:", error)
		return
	}

	var globalIp PublicIp

	error = json.Unmarshal(body, &globalIp)

	if error != nil {
		fmt.Println("Error unmarshalling initial JSON:", error)
		return
	}

	fmt.Println("Your current global ip is: ", globalIp.IP)
	return

}
