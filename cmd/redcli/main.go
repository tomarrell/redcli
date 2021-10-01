package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	cmdSubscribe = "SUBSCRIBE"
)

var rootCmd = &cobra.Command{
	Use:   "redcli",
	Short: "A simple Redis CLI for interacting with miniqueue",
}

var (
	sub   = flag.String("subscribe", "", "Subscribe to a specified topic")
	out   = flag.String("out", "", "File to write hex encoded subscribe output to")
	topic = flag.String("topic", "", "Topic to publish/subscribe to")
	host  = flag.String("host", "localhost:8080", "The hostname of the miniqueue instance")
)

func main() {
	rootCmd.AddCommand(newPublishCmd())
	rootCmd.AddCommand(newSubscribeCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
