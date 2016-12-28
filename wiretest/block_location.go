package wiretest

import (
	"math/rand"

	"github.com/TachyonNexus/alluxio-go/wire"
)

// RandomBlockLocation generates a random instance of wire.BlockLocation.
func RandomBlockLocation() wire.BlockLocation {
	return wire.BlockLocation{
		WorkerID:      rand.Int63(),
		WorkerAddress: RandomWorkerNetAddress(),
		TierAlias:     RandomString(),
	}
}
