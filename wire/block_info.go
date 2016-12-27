package wire

// BlockInfo represents a block's information.
type BlockInfo struct {
	// BlockID is the block id.
	BlockID int64
	// Length is the block length.
	Length int64
	// Locations holds the block locations.
	Locations []BlockLocation
}
