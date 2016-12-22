package wire

// ReadType represents a read type.
type ReadType string

const (
	// ReadTypeNoCache means data will be not cached.
	ReadTypeNoCache ReadType = "NO_CACHE"
	// ReadTypeCache means data will be cached.
	ReadTypeCache = "CACHE"
	// ReadTypeCachePromote mans data will be cached in the top tier.
	ReadTypeCachePromote = "CACHE_PROMOTE"
)
