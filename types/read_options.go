package types

// ReadOptions contains the possible options that can be passed to Read
type ReadOptions struct {
	ConcurrencyCount int
	BatchSize        int64
	Process          func(messages []map[string]interface{})
}
