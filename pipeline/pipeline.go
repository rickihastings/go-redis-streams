package pipeline

import (
	"github.com/rickihastings/go-redis-streams/types"
)

// Source is the interface which should be able to pull messages from a streaming
// source, and then also push to them, an example would be Redis Streams, Redis Blocking
// lists, Kafka, Kinesis, etc.
type Source interface {
	Push(interface{}, *types.Shard) error
	Read(*types.ReadOptions, types.Channel)
}

// Pipeline allows us to take an input channel, and pass the messages through
// a Processor, and then pass into another Source for processing elsewhere.
type Pipeline struct {
	input types.Channel
}

// Via defines a Processor to pass the stream through, and returns an instance
// of Pipeline so they can be changed for multiple steps.
func (p *Pipeline) Via(process types.Processor) *Pipeline {
	output := make(types.Channel)

	go func() {
		for messages := range p.input {
			output <- process(messages)
		}
	}()

	return &Pipeline{
		input: output,
	}
}

// To allows you to stream messages back into another data source for processing
// elsewhere
func (p *Pipeline) To(sink Source) {
	for messages := range p.input {
		for _, message := range messages {
			sink.Push(message.Record, message.Shard)
		}
	}
}

// Wait ensures the channel is never closed so the process runs continually
// If you're processing constant workflows you'll either want to call this or To()
func (p *Pipeline) Wait() {
	for range p.input {
		// no-op
	}
}

// From initiates a never-ending read stream from a Source, and returns
// a pipeline for processing or passing to a different Source.
func From(source Source, options *types.ReadOptions) *Pipeline {
	channel := make(types.Channel)

	go source.Read(options, channel)

	return &Pipeline{
		input: channel,
	}
}
