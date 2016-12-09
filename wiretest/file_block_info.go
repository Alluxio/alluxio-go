package wiretest

import (
	"math/rand"

	"github.com/Alluxio/alluxio-go/wire"
)

// RandomFileBlockInfo generates a random FileBlockInfo.
func RandomFileBlockInfo() wire.FileBlockInfo {
	ufsLocations := make([]string, rand.Intn(10))
	for i := 0; i < len(ufsLocations); i++ {
		ufsLocations[i] = RandomString()
	}
	return wire.FileBlockInfo{
		BlockInfo:    RandomBlockInfo(),
		Offset:       rand.Int63(),
		UfsLocations: ufsLocations,
	}
}
