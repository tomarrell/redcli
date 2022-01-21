package main

import (
	"github.com/spf13/cobra"
)

func newPublishCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "publish topic message",
		Short: "Publish a message to a topic",
		Run:   runPublish,
		Args:  cobra.ExactArgs(2),
	}
}

func runPublish(cmd *cobra.Command, args []string) {
	conn := newClient(cmd.Context())

	res, err := conn.Do("publish", args[0], args[1])
	cobra.CheckErr(err)

	cmd.Println(res)
}
