package option

type Mount struct {
	Properties map[string]string `json:"properties,omitempty"`
	ReadOnly   *bool             `json:"readOnly,omitempty"`
	Shared     *bool             `json:"shared,omitempty"`
}

func (option *Mount) SetProperties(value map[string]string) {
	option.Properties = value
}

func (option *Mount) SetReadOnly(value bool) {
	option.ReadOnly = &value
}

func (option *Mount) SetShared(value bool) {
	option.Shared = &value
}
