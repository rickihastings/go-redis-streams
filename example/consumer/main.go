package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/rickihastings/go-redis-streams/sources"
	"github.com/rickihastings/go-redis-streams/types"
)

func main() {
	stream, err := sources.NewRedisStream(&redis.Options{}, "example", "ingest")
	if err != nil {
		panic(err)
	}

	stream.Read(types.ReadOptions{
		ConcurrencyCount: 10,
		BatchSize:        1000,
	})
}
