package wire

// BlockLocation represents a block's location.
type BlockLocation struct {
	// WorkerID is the worker id.
	WorkerID int64
	// WorkerAddress is the worker network address.
	WorkerAddress WorkerNetAddress
	// TierAlias is the tier alias.
	TierAlias string
}
