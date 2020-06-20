package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/rickihastings/go-redis-streams/sources"
	"github.com/rickihastings/go-redis-streams/types"
)

type Payload struct {
	Hello string
}

// MarshalBinary is required to marshal nested objects in redis streams
func (o *Payload) MarshalBinary() (data []byte, err error) {
	return json.Marshal(o)
}

// UnmarshalBinary unmarshalls binary objects
func (o *Payload) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, o)
}

func main() {
	stream, err := sources.NewRedisStream(&redis.Options{
		Addr: os.Getenv("HOST"),
	}, "example", "ingest")
	if err != nil {
		panic(err)
	}

	record := &Payload{
		Hello: "World",
	}

	shard := &types.Shard{
		ID: "1",
	}

	wg := sync.WaitGroup{}
	semaphore := make(chan struct{}, 500)
	defer close(semaphore)

	for i := 0; i < 1506; i++ {
		go func(i int) {
			wg.Add(1)
			semaphore <- struct{}{}

			if err := stream.Push(fmt.Sprintf("1-%d", i), record, shard); err != nil {
				panic(err)
			} else {
				wg.Done()
				<-semaphore
			}
		}(i)
	}

	wg.Wait()
}
