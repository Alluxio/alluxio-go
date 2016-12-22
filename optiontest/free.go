package optiontest

import (
	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wiretest"
)

// RandomFree creates a random instance of option.Free.
func RandomFree() option.Free {
	var option option.Free
	option.SetRecursive(wiretest.RandomBool())
	return option
}
