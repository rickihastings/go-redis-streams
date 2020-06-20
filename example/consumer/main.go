package main

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/rickihastings/go-redis-streams/sources"
	"github.com/rickihastings/go-redis-streams/types"
)

func main() {
	stream, err := sources.NewRedisStream(&redis.Options{
		Addr: os.Getenv("HOST"),
	}, "example", "ingest")
	if err != nil {
		panic(err)
	}

	if err := stream.Read(types.ReadOptions{
		ConcurrencyCount: 10,
		BatchSize:        1000,
	}); err != nil {
		panic(err)
	}
}
