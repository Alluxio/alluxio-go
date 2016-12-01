package option

import (
	"github.com/Alluxio/alluxio-go/wire"
)

type CreateDirectory struct {
	AllowExists bool           `json:"allowExists"`
	Mode        *wire.Mode     `json:"mode"`
	Recursive   bool           `json:"recursive"`
	WriteType   wire.WriteType `json:"writeType"`
}

type Delete struct {
	Recursive bool `json:"recursive"`
}

type Exists struct{}

type Free struct {
	Recursive bool `json:"recursive"`
}

type GetStatus struct{}

type ListStatus struct {
	LoadMetadataType wire.LoadMetadataType `json:"loadMetadataType"`
}

type Mount struct {
	Properties map[string]string `json:"properties"`
	ReadOnly   bool              `json:"readOnly"`
	Shared     bool              `json:"shared"`
}

type Rename struct{}

type SetAttribute struct {
	Group     *string        `json:"group"`
	Mode      *wire.Mode     `json:"mode"`
	Owner     *string        `json:"owner"`
	Persisted *bool          `json:"persisted"`
	Pinned    *bool          `json:"pinned"`
	Recursive bool           `json:"recursive"`
	TTL       *int64         `json:"ttl"`
	TTLAction wire.TTLAction `json:"ttlAction"`
}

type Unmount struct{}
