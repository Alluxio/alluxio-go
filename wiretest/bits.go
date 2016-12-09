package wiretest

import (
	"math/rand"

	"github.com/Alluxio/alluxio-go/wire"
)

// RandomBits generates a random Bits.
func RandomBits() wire.Bits {
	var result wire.Bits
	switch rand.Intn(8) {
	case 0:
		result = wire.BitsNone
	case 1:
		result = wire.BitsExecute
	case 2:
		result = wire.BitsWrite
	case 3:
		result = wire.BitsWriteExecute
	case 4:
		result = wire.BitsRead
	case 5:
		result = wire.BitsReadExecute
	case 6:
		result = wire.BitsReadWrite
	case 7:
		result = wire.BitsAll
	}
	return result
}
