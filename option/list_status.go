package option

import (
	"github.com/Alluxio/alluxio-go/wire"
)

type ListStatus struct {
	LoadMetadataType *wire.LoadMetadataType `json:"loadMetadataType,omitempty"`
}

func (option *ListStatus) SetLoadMetadataType(value wire.LoadMetadataType) {
	option.LoadMetadataType = &value
}
