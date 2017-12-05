package option

import (
	"github.com/alluxio/alluxio-go/wire"
)

// OpenFile holds the options for opening a file.
type OpenFile struct {
	LocationPolicyClass *string        `json:"locationPolicyClass,omitempty"`
	ReadType            *wire.ReadType `json:"readType,omitempty"`
}

// SetLocationPolicyClass sets the location policy class.
func (option *OpenFile) SetLocationPolicyClass(value string) {
	option.LocationPolicyClass = &value
}

// SetReadType sets the read type.
func (option *OpenFile) SetReadType(value wire.ReadType) {
	option.ReadType = &value
}
