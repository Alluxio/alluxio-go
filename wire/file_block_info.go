package wire

import (
	"math/rand"
)

// FileBlockInfo represents a file block information.
type FileBlockInfo struct {
	BlockInfo    BlockInfo
	Offset       int64
	UfsLocations []string
}

// RandomFileBlockInfo generates a random file block information.
func RandomFileBlockInfo() FileBlockInfo {
	ufsLocations := make([]string, rand.Intn(10))
	for i := 0; i < len(ufsLocations); i++ {
		ufsLocations[i] = RandomString()
	}
	return FileBlockInfo{
		BlockInfo:    RandomBlockInfo(),
		Offset:       rand.Int63(),
		UfsLocations: ufsLocations,
	}
}
