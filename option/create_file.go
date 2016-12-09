package option

import "github.com/Alluxio/alluxio-go/wire"

type CreateFile struct {
	BlockSizeBytes      *int64          `json:"blockSizeBytes,omitempty"`
	LocationPolicyClass *string         `json:"locationPolicyClass,omitempty"`
	Mode                *wire.Mode      `json:"mode,omitempty"`
	Recursive           *bool           `json:"recursive,omitempty"`
	TTL                 *int64          `json:"ttl,omitempty"`
	TTLAction           *wire.TTLAction `json:"ttlAction,omitempty"`
	WriteType           *wire.WriteType `json:"writeType,omitempty"`
}

func (option *CreateFile) SetBlockSizeBytes(value int64) {
	option.BlockSizeBytes = &value
}

func (option *CreateFile) SetLocationPolicyClass(value string) {
	option.LocationPolicyClass = &value
}

func (option *CreateFile) SetMode(value wire.Mode) {
	option.Mode = &value
}

func (option *CreateFile) SetRecursive(value bool) {
	option.Recursive = &value
}

func (option *CreateFile) SetTTL(value int64) {
	option.TTL = &value
}

func (option *CreateFile) SetTTLAction(value wire.TTLAction) {
	option.TTLAction = &value
}

func (option *CreateFile) SetWriteType(value wire.WriteType) {
	option.WriteType = &value
}
