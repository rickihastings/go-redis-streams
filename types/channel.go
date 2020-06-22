package types

type Message struct {
	Metadata *Metadata
	Shard    *Shard
	Record   interface{}
}

type Channel = chan []Message
type Processor = func(id *string, messages []Message) []Message
