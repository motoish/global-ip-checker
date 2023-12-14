package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

var rate uint32

func init() {
	rootCmd.AddCommand(checkVpnCmd)
	checkVpnCmd.Flags().Uint32VarP(&rate, "rate", "r", 60, "Specify the number of seconds per interval you wish to poll")
}

var checkVpnCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch for disconnect from provided public ip",
	Run:   checkVpn,
	Args:  cobra.MinimumNArgs(1),
}

const (
	ipify   = "https://api.ipify.org?format=json"
	timeout = 10 * time.Second
)

var pollingRate = 60 * time.Second

func checkVpn(cmd *cobra.Command, args []string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	client := &http.Client{
		Timeout: timeout,
	}

	var ip string
	if validIp := net.ParseIP(args[0]); validIp == nil {
		fmt.Println("Please add a global IP for me to check")
		return
	} else {
		ip = validIp.String()
	}

	go func() {

		// Reuse the HTTP client
		response, err := client.Get(ipify)
		if err != nil {
			fmt.Println("Error initial IP check:", err)
			return
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading initial response body:", err)
			return
		}

		var initialPublicIp PublicIp
		err = json.Unmarshal(body, &initialPublicIp)
		if err != nil {
			fmt.Println("Error unmarshalling initial JSON:", err)
			return
		}

		fmt.Println(fmt.Sprintf("Watching for changes in the IP: %s", ip))

		for {
			select {
			case <-interrupt:
				fmt.Println("\nInterrupt signal received. Stopping...")
				return
			default:
				response, err := client.Get(ipify)
				if rate != 60 {
					pollingRate = time.Second * time.Duration(rate)
				}
				if err != nil {
					handleHTTPError(err)
					time.Sleep(pollingRate)
					continue
				}
				defer response.Body.Close()

				body, err := io.ReadAll(response.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					time.Sleep(pollingRate)
					continue
				}

				var publicIp PublicIp
				err = json.Unmarshal(body, &publicIp)
				if err != nil {
					fmt.Println("Error unmarshalling JSON:", err)
					time.Sleep(pollingRate)
					continue
				}

				if publicIp.IP != ip {
					handleIPChange(ip, client)
				}

				time.Sleep(pollingRate)
			}
		}
	}()

	<-interrupt
	fmt.Println("\nExiting...")
}

func handleHTTPError(err error) {
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		err := beeep.Alert("Network Error", "Something happened, you either got disconnected or your global IP changed.", "assets/warning.png")
		if err != nil {
			fmt.Println("Error sending notification:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Error getting IP:", err)
	}
}

func handleIPChange(ip string, client *http.Client) {
	client.CloseIdleConnections()
	err := beeep.Alert("Global Ip missmatched", "Your current global ip and the global ip you told me to watch doesnt match", "assets/warning.png")
	if err != nil {
		fmt.Println("Error sending notification:", err)
		os.Exit(1)
	}
}
