package wiretest

import (
	"math/rand"

	"github.com/alluxio/alluxio-go/wire"
)

// RandomTTLAction generates a random instance of wire.TTLAction.
func RandomTTLAction() wire.TTLAction {
	var result wire.TTLAction
	switch rand.Intn(2) {
	case 0:
		result = wire.TTLActionDelete
	case 1:
		result = wire.TTLActionFree
	}
	return result
}
