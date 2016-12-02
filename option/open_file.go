package option

import (
	"github.com/Alluxio/alluxio-go/wire"
)

type OpenFile struct {
	LocationPolicyClass *string        `json:"locationPolicyClass,omitempty"`
	ReadType            *wire.ReadType `json:"readType,omitempty"`
}

func (option *OpenFile) SetLocationPolicyClass(value string) {
	option.LocationPolicyClass = &value
}

func (option *OpenFile) SetReadType(value wire.ReadType) {
	option.ReadType = &value
}

func RandomOpenFile() OpenFile {
	var option OpenFile
	option.SetLocationPolicyClass(wire.RandomString())
	option.SetReadType(wire.RandomReadType())
	return option
}
