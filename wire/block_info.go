package wire

// BlockInfo represents a block's information.
type BlockInfo struct {
	BlockID   int64
	Length    int64
	Locations []BlockLocation
}
