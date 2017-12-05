package option

import (
	"github.com/alluxio/alluxio-go/wire"
)

// CreateDirectory holds the options for creating a directory.
type CreateDirectory struct {
	AllowExists *bool           `json:"allowExists,omitempty"`
	Mode        *wire.Mode      `json:"mode,omitempty"`
	Recursive   *bool           `json:"recursive,omitempty"`
	WriteType   *wire.WriteType `json:"writeType,omitempty"`
}

// SetAllowExists sets the bit that determines whether the directory to be
// created is allowed to exist.
func (option *CreateDirectory) SetAllowExists(value bool) {
	option.AllowExists = &value
}

// SetMode sets the access mode.
func (option *CreateDirectory) SetMode(value wire.Mode) {
	option.Mode = &value
}

// SetRecursive sets the bit that determines whether non-existent ancestors
// of the directory to be created should be created as well.
func (option *CreateDirectory) SetRecursive(value bool) {
	option.Recursive = &value
}

// SetWriteType sets the write type.
func (option *CreateDirectory) SetWriteType(value wire.WriteType) {
	option.WriteType = &value
}
