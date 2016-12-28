package optiontest

import (
	"github.com/TachyonNexus/alluxio-go/option"
	"github.com/TachyonNexus/alluxio-go/wiretest"
)

// RandomFree creates a random instance of option.Free.
func RandomFree() option.Free {
	var option option.Free
	option.SetRecursive(wiretest.RandomBool())
	return option
}
