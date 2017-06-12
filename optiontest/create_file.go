package optiontest

import (
	"math/rand"

	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wiretest"
)

// RandomCreateFile creates a random instance of option.CreateFile.
func RandomCreateFile() option.CreateFile {
	var option option.CreateFile
	option.SetBlockSizeBytes(rand.Int63())
	option.SetLocationPolicyClass(wiretest.RandomString())
	option.SetMode(wiretest.RandomMode())
	option.SetRecursive(wiretest.RandomBool())
	option.SetTTL(rand.Int63())
	option.SetTTLAction(wiretest.RandomTTLAction())
	option.SetWriteType(wiretest.RandomWriteType())
	return option
}
