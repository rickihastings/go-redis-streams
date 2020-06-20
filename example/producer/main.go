package main

import (
	"fmt"

	"github.com/rickihastings/go-redis-streams/types"
	"github.com/vmihailenco/msgpack"
)

type Payload struct {
	Hello string
}

// MarshalBinary is required to marshal nested objects in redis streams
func (o *Payload) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

// UnmarshalBinary unmarshalls binary objects
func (o *Payload) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}

func main() {
	// stream, err := sources.NewRedisStream(&redis.Options{
	// 	Addr: os.Getenv("HOST"),
	// }, "example", "ingest")
	// if err != nil {
	// 	panic(err)
	// }

	record := &Payload{
		Hello: "World",
	}

	shard := &types.Shard{
		ID: "1",
	}

	encoded, _ := record.MarshalBinary()
	fmt.Println(encoded)

	encoded, _ = shard.MarshalBinary()
	fmt.Println(encoded)

	panic("test")

	// for i := 0; i < 1; i++ {
	// 	if err := stream.Push(fmt.Sprintf("1-%d", i), record, shard); err != nil {
	// 		panic(err)
	// 	}
	// }
}
