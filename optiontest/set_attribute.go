package optiontest

import (
	"math/rand"

	"github.com/TachyonNexus/alluxio-go/option"
	"github.com/TachyonNexus/alluxio-go/wiretest"
)

// RandomSetAttribute creates a random instance of option.SetAttribute.
func RandomSetAttribute() option.SetAttribute {
	var option option.SetAttribute
	option.SetGroup(wiretest.RandomString())
	option.SetMode(wiretest.RandomMode())
	option.SetOwner(wiretest.RandomString())
	option.SetPersisted(wiretest.RandomBool())
	option.SetPinned(wiretest.RandomBool())
	option.SetRecursive(wiretest.RandomBool())
	option.SetTTL(rand.Int63())
	option.SetTTLAction(wiretest.RandomTTLAction())
	return option
}
