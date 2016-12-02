package option

import (
	"math/rand"

	"github.com/Alluxio/alluxio-go/wire"
)

type Mount struct {
	Properties map[string]string `json:"properties,omitempty"`
	ReadOnly   *bool             `json:"readOnly,omitempty"`
	Shared     *bool             `json:"shared,omitempty"`
}

func (option *Mount) SetProperties(value map[string]string) {
	option.Properties = value
}

func (option *Mount) SetReadOnly(value bool) {
	option.ReadOnly = &value
}

func (option *Mount) SetShared(value bool) {
	option.Shared = &value
}

func RandomMount() Mount {
	var option Mount
	properties, n := map[string]string{}, rand.Intn(10)+1
	for i := 0; i < n; i++ {
		properties[wire.RandomString()] = wire.RandomString()
	}
	option.SetProperties(properties)
	option.SetReadOnly(wire.RandomBool())
	option.SetShared(wire.RandomBool())
	return option
}
