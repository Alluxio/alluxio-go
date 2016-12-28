package optiontest

import (
	"github.com/TachyonNexus/alluxio-go/option"
	"github.com/TachyonNexus/alluxio-go/wiretest"
)

// RandomDelete creates a random instance of option.Delete.
func RandomDelete() option.Delete {
	var option option.Delete
	option.SetRecursive(wiretest.RandomBool())
	return option
}
