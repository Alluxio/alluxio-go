package wire

// WorkerNetAddress represents a worker's net address.
type WorkerNetAddress struct {
	Host     string
	RPCPort  int
	DataPort int
	WebPort  int
}
