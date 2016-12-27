package wire

// Bits represents access mode's bits.
type Bits string

const (
	// BitsNone represents no access.
	BitsNone Bits = "---"
	// BitsExecute represent execute access.
	BitsExecute = "--x"
	// BitsWrite represents write access.
	BitsWrite = "-w-"
	// BitsWriteExecute represents write and execute access.
	BitsWriteExecute = "-wx"
	// BitsRead represents read access.
	BitsRead = "r--"
	// BitsReadExecute represents read and execute access.
	BitsReadExecute = "r-x"
	// BitsReadWrite represents read and write access.
	BitsReadWrite = "rw-"
	// BitsAll represents read, write, and execute access
	BitsAll = "rwx"
)
