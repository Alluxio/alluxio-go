package wire

import "math/rand"

type LoadMetadataType string

const (
	LoadMetadataTypeNever  LoadMetadataType = "Never"
	LoadMetadataTypeOnce                    = "Once"
	LoadMetadataTypeAlways                  = "Always"
)

func RandomLoadMetadataType() LoadMetadataType {
	var result LoadMetadataType
	switch rand.Intn(3) {
	case 0:
		result = LoadMetadataTypeNever
	case 1:
		result = LoadMetadataTypeOnce
	case 2:
		result = LoadMetadataTypeAlways
	}
	return result
}
