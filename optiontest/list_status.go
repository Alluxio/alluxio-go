package optiontest

import (
	"github.com/TachyonNexus/alluxio-go/option"
	"github.com/TachyonNexus/alluxio-go/wiretest"
)

// RandomListStatus creates a random instance of option.ListStatus.
func RandomListStatus() option.ListStatus {
	var option option.ListStatus
	option.SetLoadMetadataType(wiretest.RandomLoadMetadataType())
	return option
}
