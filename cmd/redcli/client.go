package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func newClient(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	status := client.Ping(ctx)
	if err := status.Err(); err != nil {
		return nil, fmt.Errorf("pinging server: %v", err)
	}

	return client, nil
}
