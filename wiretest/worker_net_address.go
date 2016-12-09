package wiretest

import (
	"math/rand"

	"github.com/Alluxio/alluxio-go/wire"
)

// RandomWorkerNetAddress generates a random WorkerNetAddress.
func RandomWorkerNetAddress() wire.WorkerNetAddress {
	return wire.WorkerNetAddress{
		Host:     RandomString(),
		RPCPort:  rand.Int(),
		DataPort: rand.Int(),
		WebPort:  rand.Int(),
	}
}
