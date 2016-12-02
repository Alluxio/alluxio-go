package wire

import "math/rand"

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

func RandomBits() Bits {
	result, n := Bits(""), rand.Intn(8)
	if n&4 != 0 {
		result += "r"
	} else {
		result += "-"
	}
	if n&2 != 0 {
		result += "w"
	} else {
		result += "-"
	}
	if n&1 != 0 {
		result += "x"
	} else {
		result += "-"
	}
	return result
}
