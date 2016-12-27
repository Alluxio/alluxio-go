package wire

// LoadMetadataType represents the load metadata type.
type LoadMetadataType string

const (
	// LoadMetadataTypeNever means metadata should never be loaded.
	LoadMetadataTypeNever LoadMetadataType = "Never"
	// LoadMetadataTypeOnce means metadata should be loaded once.
	LoadMetadataTypeOnce = "Once"
	// LoadMetadataTypeAlways means metadata should always be loaded.
	LoadMetadataTypeAlways = "Always"
)
