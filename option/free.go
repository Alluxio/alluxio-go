package option

// Free holds the options for freeing a path.
type Free struct {
	Recursive *bool `json:"recursive,omitempty"`
}

// SetRecursive sets the bit which determines whether the path should be
// freed recursively.
func (option *Free) SetRecursive(value bool) {
	option.Recursive = &value
}
