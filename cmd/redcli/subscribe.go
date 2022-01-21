package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

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

	conn := newClient(cmd.Context())

	err := conn.Send("SUBSCRIBE", args[0])
	cobra.CheckErr(err)
	cobra.CheckErr(conn.Flush())

	for {
		res, err := redis.String(conn.Receive())
		cobra.CheckErr(err)
		fmt.Println(res)

		res, err = redis.String(conn.Do("ACK"))
		cobra.CheckErr(err)
		fmt.Println(res)
	}
}
