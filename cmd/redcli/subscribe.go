package main

import "github.com/spf13/cobra"

const subscribeCommand = "SUBSCRIBE"

var subscribeCount int

func newSubscribeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe topic",
		Short: "Subscribe to a topic",
		Run:   runSubscribe,
		Args:  cobra.ExactArgs(1),
	}

	cmd.Flags().IntVarP(&subscribeCount, "count", "c", 1, "Number of messages to consume")

	return cmd
}

func runSubscribe(cmd *cobra.Command, args []string) {
	cmd.Println("Subscribing to topic:", args[0])

	client, err := newClient(cmd.Context())

	// Make sure that our client doesn't timeout on reads
	conn := client.WithTimeout(-1).Conn(cmd.Context())
	cobra.CheckErr(err)

	conn.Process(cmd.Context()).Error()

	for i := 0; i < subscribeCount; i++ {
		res, err := conn.Do(cmd.Context(), subscribeCommand, args[0]).Result()
		cobra.CheckErr(err)

		cmd.Println(res)
	}
}
