package option

import (
	"github.com/Alluxio/alluxio-go/wire"
)

// ListStatus holds the options for listing a path.
type ListStatus struct {
	LoadMetadataType *wire.LoadMetadataType `json:"loadMetadataType,omitempty"`
}

// SetLoadMetadataType sets the type, which determines whether to load metadata
// for children of a directory.
func (option *ListStatus) SetLoadMetadataType(value wire.LoadMetadataType) {
	option.LoadMetadataType = &value
}
