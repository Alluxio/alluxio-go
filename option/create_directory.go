package option

import (
	"github.com/Alluxio/alluxio-go/wire"
)

type CreateDirectory struct {
	AllowExists *bool           `json:"allowExists,omitempty"`
	Mode        *wire.Mode      `json:"mode,omitempty"`
	Recursive   *bool           `json:"recursive,omitempty"`
	WriteType   *wire.WriteType `json:"writeType,omitempty"`
}

func (option *CreateDirectory) SetAllowExists(value bool) {
	option.AllowExists = &value
}

func (option *CreateDirectory) SetMode(value wire.Mode) {
	option.Mode = &value
}

func (option *CreateDirectory) SetRecursive(value bool) {
	option.Recursive = &value
}

func (option *CreateDirectory) SetWriteType(value wire.WriteType) {
	option.WriteType = &value
}

func RandomCreateDirectory() CreateDirectory {
	var option CreateDirectory
	option.SetAllowExists(wire.RandomBool())
	option.SetMode(wire.RandomMode())
	option.SetRecursive(wire.RandomBool())
	option.SetWriteType(wire.RandomWriteType())
	return option
}
