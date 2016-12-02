package option

import "github.com/Alluxio/alluxio-go/wire"

type Free struct {
	Recursive *bool `json:"recursive,omitempty"`
}

func (option *Free) SetRecursive(value bool) {
	option.Recursive = &value
}

func RandomFree() Free {
	var option Free
	option.SetRecursive(wire.RandomBool())
	return option
}
