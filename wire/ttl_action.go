package wire

// TTLAction represents a TTL action.
type TTLAction string

const (
	TTLActionDelete TTLAction = "DELETE"
	TTLActionFree             = "FREE"
)
