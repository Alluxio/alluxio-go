package wire

// FileBlockInfo represents a file block's information.
type FileBlockInfo struct {
	BlockInfo    BlockInfo
	Offset       int64
	UfsLocations []string
}
