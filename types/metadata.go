package types

import "github.com/vmihailenco/msgpack/v5"

// Metadata is a type used for message metadata
type Metadata struct{}

// MarshalBinary is required to marshal nested objects in redis streams
func (o *Metadata) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

// UnmarshalBinary unmarshalls binary objects
func (o *Metadata) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
