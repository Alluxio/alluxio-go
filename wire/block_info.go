package wire

import (
	"math/rand"
)

// BlockInfo represents a block information.
type BlockInfo struct {
	BlockID   int64
	Length    int64
	Locations []BlockLocation
}

// RandomBlockInfo generates a random block information.
func RandomBlockInfo() BlockInfo {
	locations := make([]BlockLocation, rand.Intn(10))
	for i := 0; i < len(locations); i++ {
		locations[i] = RandomBlockLocation()
	}
	return BlockInfo{
		BlockID:   rand.Int63(),
		Length:    rand.Int63(),
		Locations: locations,
	}
}
