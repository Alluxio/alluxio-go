package wiretest

import (
	"math/rand"

	"github.com/alluxio/alluxio-go/wire"
)

// RandomWorkerNetAddress generates a random instance of wire.WorkerNetAddress.
func RandomWorkerNetAddress() wire.WorkerNetAddress {
	return wire.WorkerNetAddress{
		Host:     RandomString(),
		RPCPort:  rand.Int(),
		DataPort: rand.Int(),
		WebPort:  rand.Int(),
	}
}
