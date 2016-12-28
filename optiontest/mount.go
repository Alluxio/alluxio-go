package optiontest

import (
	"math/rand"

	"github.com/TachyonNexus/alluxio-go/option"
	"github.com/TachyonNexus/alluxio-go/wiretest"
)

// RandomMount creates a random instance of option.Mount.
func RandomMount() option.Mount {
	var option option.Mount
	properties, n := map[string]string{}, rand.Intn(10)+1
	for i := 0; i < n; i++ {
		properties[wiretest.RandomString()] = wiretest.RandomString()
	}
	option.SetProperties(properties)
	option.SetReadOnly(wiretest.RandomBool())
	option.SetShared(wiretest.RandomBool())
	return option
}
