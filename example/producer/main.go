package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/rickihastings/go-redis-streams/sources"
	"github.com/rickihastings/go-redis-streams/types"
)

func main() {
	stream, err := sources.NewRedisStream(&redis.Options{}, "example", "ingest")
	if err != nil {
		panic(err)
	}

	record := types.Record{
		"Hello": "World",
	}

	shard := types.Shard{
		ID: "1",
	}

	for i := 0; i < 10; i++ {
		stream.Push(fmt.Sprintf("1-%d", i), record, shard)
	}
}
