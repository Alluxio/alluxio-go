package wiretest

import (
	"math/rand"

	"github.com/alluxio/alluxio-go/wire"
)

// RandomReadType generates a random instance of wire.ReadType.
func RandomReadType() wire.ReadType {
	var result wire.ReadType
	switch rand.Intn(3) {
	case 0:
		result = wire.ReadTypeNoCache
	case 1:
		result = wire.ReadTypeCache
	case 2:
		result = wire.ReadTypeCachePromote
	}
	return result
}
