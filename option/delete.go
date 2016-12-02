package option

type Delete struct {
	Recursive *bool `json:"recursive,omitempty"`
}

func (option *Delete) SetRecursive(value bool) {
	option.Recursive = &value
}
