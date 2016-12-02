package option

import "github.com/Alluxio/alluxio-go/wire"

type Delete struct {
	Recursive *bool `json:"recursive,omitempty"`
}

func (option *Delete) SetRecursive(value bool) {
	option.Recursive = &value
}

func RandomDelete() Delete {
	var option Delete
	option.SetRecursive(wire.RandomBool())
	return option
}
