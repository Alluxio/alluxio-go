package option

import "github.com/Alluxio/alluxio-go/wire"

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

func (option *SetAttribute) SetGroup(value string) {
	option.Group = &value
}

func (option *SetAttribute) SetMode(value wire.Mode) {
	option.Mode = &value
}

func (option *SetAttribute) SetOwner(value string) {
	option.Owner = &value
}

func (option *SetAttribute) SetPersisted(value bool) {
	option.Persisted = &value
}

func (option *SetAttribute) SetPinned(value bool) {
	option.Pinned = &value
}

func (option *SetAttribute) SetRecursive(value bool) {
	option.Recursive = &value
}

func (option *SetAttribute) SetTTL(value int64) {
	option.TTL = &value
}

func (option *SetAttribute) SetTTLAction(value wire.TTLAction) {
	option.TTLAction = &value
}
