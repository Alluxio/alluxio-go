package wire

type Mode struct {
	OwnerBits Bits `json:"ownerBits"`
	GroupBits Bits `json:"groupBits"`
	OtherBits Bits `json:"otherBits"`
}

func RandomMode() Mode {
	return Mode{
		OwnerBits: RandomBits(),
		GroupBits: RandomBits(),
		OtherBits: RandomBits(),
	}
}
