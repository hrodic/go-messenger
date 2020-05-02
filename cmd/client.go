package cmd

import (
	"strconv"
	"fmt"
	"net"
	"github.com/spf13/cobra"
)

const (
	maxBufferSize = 1024
)

var port uint
var message string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Messenger client",
	Long: `Messenger client`,
	Args: cobra.OnlyValidArgs,
	ValidArgs: []string{"port", "message"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Client connecting to: \"%s\"\n", strconv.Itoa(int(port)))

		conn, _ := net.Dial("udp", ":"+strconv.Itoa(int(port)))
		defer conn.Close()
		
		request := []byte(message)
		fmt.Printf("Sending request: \"%s\"\n", request)
		conn.Write(request)
		response := make([]byte, maxBufferSize)
		conn.Read(response)
		fmt.Printf("Received response: \"%s\"\n", response)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// local flags
	clientCmd.Flags().UintVar(&port, "p", 8888, "Port number")
	clientCmd.Flags().StringVar(&message, "m", "Hello from client", "The message")
}
