package wire

type Mode struct {
	OwnerBits Bits `json:"ownerBits"`
	GroupBits Bits `json:"groupBits"`
	OtherBits Bits `json:"otherBits"`
}
