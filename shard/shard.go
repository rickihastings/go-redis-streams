package shard

import "github.com/rickihastings/go-redis-streams/types"

// Shards is a map of all the shards identified in the block of messages
// with a list of all the messages as the value
type Shards = map[string][]types.Message

// GroupBy will take a list of messages and group them by the shard
func GroupBy(messages []types.Message) Shards {
	groups := make(Shards)

	for _, message := range messages {
		if groups[message.Shard.ID] == nil {
			groups[message.Shard.ID] = []types.Message{}
		}

		groups[message.Shard.ID] = append(groups[message.Shard.ID], message)
	}

	return groups
}
