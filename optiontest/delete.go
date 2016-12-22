package optiontest

import (
	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wiretest"
)

// RandomDelete creates a random instance of option.Delete.
func RandomDelete() option.Delete {
	var option option.Delete
	option.SetRecursive(wiretest.RandomBool())
	return option
}
