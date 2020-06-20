package types

import "encoding/json"

// Metadata is a type used for message metadata
type Metadata struct{}

// MarshalBinary is required to marshal nested objects in redis streams
func (o *Metadata) MarshalBinary() (data []byte, err error) {
	return json.Marshal(o)
}

// UnmarshalBinary unmarshalls binary objects
func (o *Metadata) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, o)
}
