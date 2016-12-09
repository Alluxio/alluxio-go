package optiontest

import (
	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wiretest"
)

func RandomCreateDirectory() option.CreateDirectory {
	var option option.CreateDirectory
	option.SetAllowExists(wiretest.RandomBool())
	option.SetMode(wiretest.RandomMode())
	option.SetRecursive(wiretest.RandomBool())
	option.SetWriteType(wiretest.RandomWriteType())
	return option
}
