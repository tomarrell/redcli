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
	client, err := newClient(cmd.Context())
	cobra.CheckErr(err)

	res, err := client.Do(cmd.Context(), "publish", args[0], args[1]).Result()
	cobra.CheckErr(err)

	cmd.Println(res)
}
