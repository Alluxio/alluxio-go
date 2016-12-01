package wire

type LoadMetadataType string

const (
	LoadMetadataTypeNever  LoadMetadataType = "Never"
	LoadMetadataTypeOnce                    = "Once"
	LoadMetadataTypeAlways                  = "Always"
)
