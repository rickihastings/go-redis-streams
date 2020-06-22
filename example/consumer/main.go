package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rickihastings/go-redis-streams/pipeline"
	"github.com/rickihastings/go-redis-streams/sources"
	"github.com/rickihastings/go-redis-streams/types"
)

func process(messages []types.Message) []types.Message {
	for _, msg := range messages {
		fmt.Println("Consumer1:", msg)
	}

	return messages
}

func main() {
	// Wait until redis is ready
	time.Sleep(1 * time.Second)

	stream, err := sources.NewRedisStream(&redis.Options{
		Addr: "redis:6379",
	}, "example", "ingest")
	if err != nil {
		panic(err)
	}

	pipeline.
		From(stream, &types.ReadOptions{
			ConcurrencyCount: 10,
			BatchSize:        100,
		}).
		Via(process).
		Wait()
}
