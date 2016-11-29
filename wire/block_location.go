package wire

import (
	"math/rand"
)

// BlockLocation represents a block location.
type BlockLocation struct {
	WorkerID      int64
	WorkerAddress WorkerNetAddress
	TierAlias     string
}

// RandomBlockLocation generates a random block location.
func RandomBlockLocation() BlockLocation {
	return BlockLocation{
		WorkerID:      rand.Int63(),
		WorkerAddress: RandomWorkerNetAddress(),
		TierAlias:     RandomString(),
	}
}
