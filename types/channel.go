package types

type Message = map[string]interface{}
type Channel = chan []Message

type Processor = func(messages []Message) []Message
