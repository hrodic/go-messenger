package cmd

import (
	"fmt"
	"strings"
	"context"
	"flag"
	"github.com/hrodic/go-messenger/internal/udp"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "HTTP Server",
	Long: `Serve messenger via HTTP`,
	Args: cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))

		address := flag.String("address", ":8888",
		"The UDP Server listen address with port, e.g. `:8888` or `0.0.0.0:8888`.")
		flag.Parse()
		ctx, _ := context.WithCancel(context.Background())
		
		go udp.Serve(ctx, *address)
		fmt.Printf("UDP Server listening on: \"%s\"\n", *address)
		<-ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}