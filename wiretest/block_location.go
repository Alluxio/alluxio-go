package wiretest

import (
	"math/rand"

	"github.com/Alluxio/alluxio-go/wire"
)

// RandomBlockLocation generates a random BlockLocation.
func RandomBlockLocation() wire.BlockLocation {
	return wire.BlockLocation{
		WorkerID:      rand.Int63(),
		WorkerAddress: RandomWorkerNetAddress(),
		TierAlias:     RandomString(),
	}
}
