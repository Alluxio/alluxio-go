package wire

// WorkerNetAddress represents a worker's net address.
type WorkerNetAddress struct {
	// Host is the hostname.
	Host string
	// RPCPort is the RPC port.
	RPCPort int
	// DataPort is the data port.
	DataPort int
	// WebPort is the web port.
	WebPort int
}
