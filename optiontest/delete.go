package optiontest

import (
	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wiretest"
)

func RandomDelete() option.Delete {
	var option option.Delete
	option.SetRecursive(wiretest.RandomBool())
	return option
}
