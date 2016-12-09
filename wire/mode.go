package wire

// Mode represents the file's mode.
type Mode struct {
	OwnerBits Bits `json:"ownerBits"`
	GroupBits Bits `json:"groupBits"`
	OtherBits Bits `json:"otherBits"`
}
