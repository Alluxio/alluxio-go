package wire

import (
	"math/rand"
)

// WorkerNetAddress represents a worker net address.
type WorkerNetAddress struct {
	Host     string
	RPCPort  int
	DataPort int
	WebPort  int
}

// RandomWorkerNetAddress generates a random worker net address.
func RandomWorkerNetAddress() WorkerNetAddress {
	return WorkerNetAddress{
		Host:     RandomString(),
		RPCPort:  rand.Int(),
		DataPort: rand.Int(),
		WebPort:  rand.Int(),
	}
}
