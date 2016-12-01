package wire

type WriteType string

const (
	WriteTypeMustCache    WriteType = "MUST_CACHE"
	WriteTypeCacheThrough           = "CACHE_THROUGH"
	WriteTypeThrough                = "THROUGH"
	WriteTypeAsyncThrough           = "ASYNC_THROUGH"
	WriteTypeNone                   = "NONE"
)
