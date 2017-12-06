package option

import "github.com/alluxio/alluxio-go/wire"

// SetAttribute holds the options for setting path attributes.
type SetAttribute struct {
	Group     *string         `json:"group,omitempty"`
	Mode      *wire.Mode      `json:"mode,omitempty"`
	Owner     *string         `json:"owner,omitempty"`
	Persisted *bool           `json:"persisted,omitempty"`
	Pinned    *bool           `json:"pinned,omitempty"`
	Recursive *bool           `json:"recursive,omitempty"`
	TTL       *int64          `json:"ttl,omitempty"`
	TTLAction *wire.TTLAction `json:"ttlAction,omitempty"`
}

// SetGroup sets the group.
func (option *SetAttribute) SetGroup(value string) {
	option.Group = &value
}

// SetMode sets the access mode.
func (option *SetAttribute) SetMode(value wire.Mode) {
	option.Mode = &value
}

// SetOwner sets the owner.
func (option *SetAttribute) SetOwner(value string) {
	option.Owner = &value
}

// SetPersisted sets the bit which determines whether the path is persisted.
func (option *SetAttribute) SetPersisted(value bool) {
	option.Persisted = &value
}

// SetPinned sets the bit which determines whether a path is pinned.
func (option *SetAttribute) SetPinned(value bool) {
	option.Pinned = &value
}

// SetRecursive sets the bit which determines whether the attributes should
// be set recursively.
func (option *SetAttribute) SetRecursive(value bool) {
	option.Recursive = &value
}

// SetTTL sets the time-to-live window. Once this window expires a time-to-live
// action is performed.
func (option *SetAttribute) SetTTL(value int64) {
	option.TTL = &value
}

// SetTTLAction sets the time-to-live action, which is execute once the
// time-to-live window expires.
func (option *SetAttribute) SetTTLAction(value wire.TTLAction) {
	option.TTLAction = &value
}
