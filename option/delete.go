package option

// Delete holds the options for deleting a path.
type Delete struct {
	Recursive *bool `json:"recursive,omitempty"`
}

// SetRecursive sets the bit which determines whether the path should be
// deleted recursively.
func (option *Delete) SetRecursive(value bool) {
	option.Recursive = &value
}
