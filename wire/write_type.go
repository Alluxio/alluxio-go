package wire

import "math/rand"

type WriteType string

const (
	WriteTypeMustCache    WriteType = "MUST_CACHE"
	WriteTypeCacheThrough           = "CACHE_THROUGH"
	WriteTypeThrough                = "THROUGH"
	WriteTypeAsyncThrough           = "ASYNC_THROUGH"
	WriteTypeNone                   = "NONE"
)

func RandomWriteType() WriteType {
	var result WriteType
	switch rand.Intn(5) {
	case 0:
		result = WriteTypeMustCache
	case 1:
		result = WriteTypeCacheThrough
	case 2:
		result = WriteTypeThrough
	case 3:
		result = WriteTypeAsyncThrough
	case 4:
		result = WriteTypeNone
	}
	return result
}
