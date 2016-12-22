package wire

import "sort"

// FileInfo represents a file's information.
type FileInfo struct {
	// BlockIds holds the block ids.
	BlockIds []int64
	// BlockSizeBytes is the block size (in bytes).
	BlockSizeBytes int64
	// Cacheable determines whether the file is cacheable.
	Cacheable bool
	// Completed determines whether the file is completed.
	Completed bool
	// CreationTimesMs is the creation time (in milliseconds).
	CreationTimeMs int64
	// FileBlockInfos holds the file block information.
	FileBlockInfos []FileBlockInfo
	// FileID is the file id.
	FileID int64
	// Folder determines whether the file is a folder.
	Folder bool
	// Group is the group.
	Group string
	// InMemoryPercentage represents the in-memory percentage.
	InMemoryPercentage int32
	// LastModificationTimeMs is the last modification time (in milliseconds).
	LastModificationTimeMs int64
	// Length is the file length.
	Length int64
	// Name is the file name.
	Name string
	// Path is the file path.
	Path string
	// Persisted determines whether file is persisted.
	Persisted bool
	// PersistenceState represents the persistence state.
	PersistenceState string
	// Pinned determines whether the file is pinned.
	Pinned bool
	// Mode is the access mode.
	Mode int32
	// MountPoint determines whether the file is a mount point.
	MountPoint bool
	// Owner is the owner.
	Owner string
	// TTL is the time-to-live window.
	TTL int64
	// TTLAction si the time-to-live action.
	TTLAction string
	// UfsPath is the UFS path.
	UfsPath string
}

// FileInfos represents a list of file information.
type FileInfos []FileInfo

var _ sort.Interface = (*FileInfos)(nil)

func (fis FileInfos) Len() int {
	return len(fis)
}

func (fis FileInfos) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis FileInfos) Less(i, j int) bool {
	return fis[i].Name < fis[j].Name
}
