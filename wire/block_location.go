package wire

// BlockLocation represents a block's location.
type BlockLocation struct {
	WorkerID      int64
	WorkerAddress WorkerNetAddress
	TierAlias     string
}
