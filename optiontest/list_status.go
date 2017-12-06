package optiontest

import (
	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wiretest"
)

// RandomListStatus creates a random instance of option.ListStatus.
func RandomListStatus() option.ListStatus {
	var option option.ListStatus
	option.SetLoadMetadataType(wiretest.RandomLoadMetadataType())
	return option
}
