package wire

// TTLAction represents a TTL action.
type TTLAction string

const (
	// TTLActionDelete represents the action of deleting a path.
	TTLActionDelete TTLAction = "DELETE"
	// TTLActionFree represents the action of freeing a path.
	TTLActionFree = "FREE"
)
