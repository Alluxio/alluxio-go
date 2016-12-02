package wire

import "math/rand"

type TTLAction string

const (
	TTLActionDelete TTLAction = "DELETE"
	TTLActionFree             = "FREE"
)

func RandomTTLAction() TTLAction {
	var result TTLAction
	switch rand.Intn(2) {
	case 0:
		result = TTLActionDelete
	case 1:
		result = TTLActionFree
	}
	return result
}
