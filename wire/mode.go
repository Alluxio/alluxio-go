package wire

// Mode represents the file's access mode.
type Mode struct {
	// OwnerBits represents the owner access mode
	OwnerBits Bits `json:"ownerBits"`
	// GroupBits represents the group access mode
	GroupBits Bits `json:"groupBits"`
	// OtherBits represents the other access mode
	OtherBits Bits `json:"otherBits"`
}
