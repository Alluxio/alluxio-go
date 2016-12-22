package option

// Mount holds the options for mounting a path.
type Mount struct {
	Properties map[string]string `json:"properties,omitempty"`
	ReadOnly   *bool             `json:"readOnly,omitempty"`
	Shared     *bool             `json:"shared,omitempty"`
}

// SetProperties sets the mount point properties.
func (option *Mount) SetProperties(value map[string]string) {
	option.Properties = value
}

// SetReadOnly sets the bit which determines whether the mount point is
// read-only.
func (option *Mount) SetReadOnly(value bool) {
	option.ReadOnly = &value
}

// SetShared which determines whether this mount point is shared among multiple
// users.
func (option *Mount) SetShared(value bool) {
	option.Shared = &value
}
