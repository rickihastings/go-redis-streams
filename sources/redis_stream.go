package sources

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/rickihastings/go-redis-streams/types"
)

// RedisStream implements Redis Streams
type RedisStream struct {
	name    string
	step    string
	redisdb *redis.Client
	ctx     context.Context
}

// NewRedisStream implements a new RedisStream instance
func NewRedisStream(config *redis.Options, name, step string) (*RedisStream, error) {
	ctx, _ := context.WithCancel(context.Background())

	redisdb := redis.NewClient(config)

	_, err := redisdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	source := &RedisStream{
		name,
		step,
		redisdb,
		ctx,
	}

	return source, nil
}

// Push will push a record to the stream
func (r *RedisStream) Push(id string, record interface{}, shard *types.Shard) error {
	return r.redisdb.XAdd(r.ctx, &redis.XAddArgs{
		Stream: r.name,
		Values: types.Message{
			"metadata": &types.Metadata{},
			"shard":    shard,
			"record":   record,
		},
	}).Err()
}

// Read starts to read from ConsumerCount at rate of BatchSize
func (r *RedisStream) Read(options *types.ReadOptions, channel types.Channel) {
	// Set some sensible defaults
	if options.ConcurrencyCount == 0 {
		options.ConcurrencyCount = 1
	}

	if options.BatchSize == 0 {
		options.BatchSize = 1000
	}

	// Setup the stream for reading, we need a consumer group
	r.redisdb.XGroupCreateMkStream(r.ctx, r.name, r.step, "$")

	// Create a waitgroup of concurrency size
	i := 0
	wg := sync.WaitGroup{}
	semaphore := make(chan struct{}, options.ConcurrencyCount)
	defer close(semaphore)

	for {
		// Increment the wait group abd block
		i++
		wg.Add(1)
		semaphore <- struct{}{}

		go func(i int) {
			err := r.Consume(i, options.BatchSize, channel)

			// Panic on an error - not sure what you'd do here, but it's possible to
			// get stuck in a loop of spewing errors out which isn't ideal.
			if err != nil {
				panic(err)
			}

			wg.Done()
			<-semaphore
		}(i)
	}
}

// Consume will consume messages from the stream
func (r *RedisStream) Consume(i int, batchSize int64, channel types.Channel) error {
	res, err := r.redisdb.XReadGroup(r.ctx, &redis.XReadGroupArgs{
		Group:    r.step,
		Consumer: fmt.Sprintf("consumer-%d", i),
		Streams:  []string{r.name, ">"},
		Count:    batchSize,
	}).Result()

	if err != nil {
		return err
	}

	for _, stream := range res {
		batch := make([]types.Message, 0)

		for _, message := range stream.Messages {

			metadata := &types.Metadata{}
			if rawMetadata, ok := message.Values["metadata"].(string); ok {
				if err := metadata.UnmarshalBinary([]byte(rawMetadata)); err != nil {
					return err
				}
			}

			shard := &types.Shard{}
			if rawShard, ok := message.Values["shard"].(string); ok {
				if err := shard.UnmarshalBinary([]byte(rawShard)); err != nil {
					return err
				}
			}

			msg := types.Message{
				"metadata": metadata,
				"shard":    shard,
				"record":   message.Values["record"],
			}

			r.redisdb.XAck(r.ctx, r.name, r.step, message.ID)

			batch = append(batch, msg)
		}

		channel <- batch
	}

	return nil
}
