package wiretest

import "github.com/Alluxio/alluxio-go/wire"

// RandomMode generates a random Mode.
func RandomMode() wire.Mode {
	return wire.Mode{
		OwnerBits: RandomBits(),
		GroupBits: RandomBits(),
		OtherBits: RandomBits(),
	}
}
