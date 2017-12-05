package option

import "github.com/alluxio/alluxio-go/wire"

// CreateFile holds the options for creating a file.
type CreateFile struct {
	BlockSizeBytes      *int64          `json:"blockSizeBytes,omitempty"`
	LocationPolicyClass *string         `json:"locationPolicyClass,omitempty"`
	Mode                *wire.Mode      `json:"mode,omitempty"`
	Recursive           *bool           `json:"recursive,omitempty"`
	TTL                 *int64          `json:"ttl,omitempty"`
	TTLAction           *wire.TTLAction `json:"ttlAction,omitempty"`
	WriteType           *wire.WriteType `json:"writeType,omitempty"`
}

// SetBlockSizeBytes sets the block size (in bytes).
func (option *CreateFile) SetBlockSizeBytes(value int64) {
	option.BlockSizeBytes = &value
}

// SetLocationPolicyClass sets the location policy class.
func (option *CreateFile) SetLocationPolicyClass(value string) {
	option.LocationPolicyClass = &value
}

// SetMode sets the mode.
func (option *CreateFile) SetMode(value wire.Mode) {
	option.Mode = &value
}

// SetRecursive sets the bit that determines whether non-existent ancestors
// of the file to be created should be created as well.
func (option *CreateFile) SetRecursive(value bool) {
	option.Recursive = &value
}

// SetTTL sets the time-to-live window. Once this window expires a time-to-live
// action is performed.
func (option *CreateFile) SetTTL(value int64) {
	option.TTL = &value
}

// SetTTLAction sets the time-to-live action, which is execute once the
// time-to-live window expires.
func (option *CreateFile) SetTTLAction(value wire.TTLAction) {
	option.TTLAction = &value
}

// SetWriteType sets the write type.
func (option *CreateFile) SetWriteType(value wire.WriteType) {
	option.WriteType = &value
}
