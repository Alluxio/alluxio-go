package wire

// Bits represents access mode's bits.
type Bits string

const (
	// BitsNone represents no access.
	BitsNone Bits = "NONE"
	// BitsExecute represent execute access.
	BitsExecute = "EXECUTE"
	// BitsWrite represents write access.
	BitsWrite = "WRITE"
	// BitsWriteExecute represents write and execute access.
	BitsWriteExecute = "WRITEEXECUTE"
	// BitsRead represents read access.
	BitsRead = "READ"
	// BitsReadExecute represents read and execute access.
	BitsReadExecute = "READEXECUTE"
	// BitsReadWrite represents read and write access.
	BitsReadWrite = "READWRITE"
	// BitsAll represents read, write, and execute access
	BitsAll = "ALL"
)
