package main

import (
	"context"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
)

func newClient(ctx context.Context) redis.Conn {
	conn, err := redis.DialURL("redis://localhost:6379")
	cobra.CheckErr(err)

	_, err = conn.Do("ping")
	cobra.CheckErr(err)

	return conn
}
