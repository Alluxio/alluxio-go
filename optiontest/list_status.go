package optiontest

import (
	"github.com/alluxio/alluxio-go/option"
	"github.com/alluxio/alluxio-go/wiretest"
)

// RandomListStatus creates a random instance of option.ListStatus.
func RandomListStatus() option.ListStatus {
	var option option.ListStatus
	option.SetLoadMetadataType(wiretest.RandomLoadMetadataType())
	return option
}
