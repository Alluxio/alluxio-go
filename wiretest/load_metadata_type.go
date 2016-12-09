package wiretest

import (
	"math/rand"

	"github.com/Alluxio/alluxio-go/wire"
)

// RandomLoadMetadataType generates a random LoadMetadataType.
func RandomLoadMetadataType() wire.LoadMetadataType {
	var result wire.LoadMetadataType
	switch rand.Intn(3) {
	case 0:
		result = wire.LoadMetadataTypeNever
	case 1:
		result = wire.LoadMetadataTypeOnce
	case 2:
		result = wire.LoadMetadataTypeAlways
	}
	return result
}
