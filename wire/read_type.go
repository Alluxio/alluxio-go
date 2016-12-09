package wire

// ReadType represents a read type.
type ReadType string

const (
	ReadTypeNoCache      ReadType = "NO_CACHE"
	ReadTypeCache                 = "CACHE"
	ReadTypeCachePromote          = "CACHE_PROMOTE"
)
