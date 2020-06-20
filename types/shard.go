package types

import (
	"github.com/vmihailenco/msgpack/v5"
)

// Shard is a type used to group elements from the same set together
type Shard struct {
	ID string
}

// MarshalBinary is required to marshal nested objects in redis streams
func (o *Shard) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

// UnmarshalBinary unmarshalls binary objects
func (o *Shard) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
