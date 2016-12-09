package optiontest

import (
	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wiretest"
)

func RandomListStatus() option.ListStatus {
	var option option.ListStatus
	option.SetLoadMetadataType(wiretest.RandomLoadMetadataType())
	return option
}
