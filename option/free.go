package option

type Free struct {
	Recursive *bool `json:"recursive,omitempty"`
}

func (option *Free) SetRecursive(value bool) {
	option.Recursive = &value
}
