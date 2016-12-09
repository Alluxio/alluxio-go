package wire

// LoadMetadataType represents the load metadata type.
type LoadMetadataType string

const (
	LoadMetadataTypeNever  LoadMetadataType = "Never"
	LoadMetadataTypeOnce                    = "Once"
	LoadMetadataTypeAlways                  = "Always"
)
