package wire

type Bits string

const (
	BitsNone         Bits = "---"
	BitsExecute           = "--x"
	BitsWrite             = "-w-"
	BitsWriteExecute      = "-wx"
	BitsRead              = "r--"
	BitsReadExecute       = "r-x"
	BitsReadWrite         = "rw-"
	BitsAll               = "r--"
)
