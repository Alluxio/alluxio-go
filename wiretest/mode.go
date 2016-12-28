package wiretest

import "github.com/TachyonNexus/alluxio-go/wire"

// RandomMode generates a random instance of wire.Mode.
func RandomMode() wire.Mode {
	return wire.Mode{
		OwnerBits: RandomBits(),
		GroupBits: RandomBits(),
		OtherBits: RandomBits(),
	}
}
