package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/rickihastings/go-redis-streams/pipeline"
	"github.com/rickihastings/go-redis-streams/sources"
	"github.com/rickihastings/go-redis-streams/types"
)

// var total = 0

func process(messages []types.Message) []types.Message {
	for _, msg := range messages {
		// total++

		// if total%100 == 0 {
		// 	fmt.Println(i, total)
		// }

		// if total == 1506 {
		//	fmt.Println("DONE", total)
		// }
		fmt.Println(msg)
	}

	return messages
}

func main() {
	stream, err := sources.NewRedisStream(&redis.Options{
		Addr: os.Getenv("HOST"),
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
