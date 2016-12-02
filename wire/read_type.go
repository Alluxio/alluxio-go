package wire

import "math/rand"

type ReadType string

const (
	ReadTypeNoCache      ReadType = "NO_CACHE"
	ReadTypeCache                 = "CACHE"
	ReadTypeCachePromote          = "CACHE_PROMOTE"
)

func RandomReadType() ReadType {
	var result ReadType
	switch rand.Intn(3) {
	case 0:
		result = ReadTypeNoCache
	case 1:
		result = ReadTypeCache
	case 2:
		result = ReadTypeCachePromote
	}
	return result
}
