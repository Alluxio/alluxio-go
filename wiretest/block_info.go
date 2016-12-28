package wiretest

import (
	"math/rand"

	"github.com/TachyonNexus/alluxio-go/wire"
)

// RandomBlockInfo generates a random instance of wire.BlockInfo.
func RandomBlockInfo() wire.BlockInfo {
	locations := make([]wire.BlockLocation, rand.Intn(10))
	for i := 0; i < len(locations); i++ {
		locations[i] = RandomBlockLocation()
	}
	return wire.BlockInfo{
		BlockID:   rand.Int63(),
		Length:    rand.Int63(),
		Locations: locations,
	}
}
