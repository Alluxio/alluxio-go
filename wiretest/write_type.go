package wiretest

import (
	"math/rand"

	"github.com/alluxio/alluxio-go/wire"
)

// RandomWriteType generates a random instance of wire.WriteType.
func RandomWriteType() wire.WriteType {
	var result wire.WriteType
	switch rand.Intn(5) {
	case 0:
		result = wire.WriteTypeMustCache
	case 1:
		result = wire.WriteTypeCacheThrough
	case 2:
		result = wire.WriteTypeThrough
	case 3:
		result = wire.WriteTypeAsyncThrough
	case 4:
		result = wire.WriteTypeNone
	}
	return result
}
