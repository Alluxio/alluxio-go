package optiontest

import (
	"github.com/TachyonNexus/alluxio-go/option"
	"github.com/TachyonNexus/alluxio-go/wiretest"
)

// RandomCreateDirectory creates a random instance of option.CreateDirectory.
func RandomCreateDirectory() option.CreateDirectory {
	var option option.CreateDirectory
	option.SetAllowExists(wiretest.RandomBool())
	option.SetMode(wiretest.RandomMode())
	option.SetRecursive(wiretest.RandomBool())
	option.SetWriteType(wiretest.RandomWriteType())
	return option
}
