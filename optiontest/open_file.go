package optiontest

import (
	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wiretest"
)

// RandomOpenFile creates a random instance of option.OpenFile.
func RandomOpenFile() option.OpenFile {
	var option option.OpenFile
	option.SetLocationPolicyClass(wiretest.RandomString())
	option.SetReadType(wiretest.RandomReadType())
	return option
}
