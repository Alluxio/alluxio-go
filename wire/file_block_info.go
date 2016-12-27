package wire

// FileBlockInfo represents a file block's information.
type FileBlockInfo struct {
	// BlockInfo is the block information
	BlockInfo BlockInfo
	// Offset is the file offset.
	Offset int64
	// UfsLocations holds the UFS locations.
	UfsLocations []string
}
