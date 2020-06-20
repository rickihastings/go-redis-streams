package sources

import "github.com/go-redis/redis"

// RedisStream implements Redis Streams
type RedisStream struct {
	redisdb *redis.Client
}

// NewRedisStream implements a new RedisStream instance
func NewRedisStream(config *redis.Options) *RedisStream, error {
	redisdb := redis.NewClient(config)

	source := &RedisStream{
		redisdb,
	}

	return source, nil
}
