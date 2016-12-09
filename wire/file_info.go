package wire

import "sort"

// FileInfo represents a file's information.
type FileInfo struct {
	BlockIds               []int64
	Cacheable              bool
	Completed              bool
	CreationTimeMs         int64
	BlockSizeBytes         int64
	FileBlockInfos         []FileBlockInfo
	FileID                 int64
	Folder                 bool
	Group                  string
	InMemoryPercentage     int32
	LastModificationTimeMs int64
	Length                 int64
	Name                   string
	Path                   string
	Persisted              bool
	Pinned                 bool
	Mode                   int32
	MountPoint             bool
	Owner                  string
	PersistenceState       string
	TTL                    int64
	TTLAction              string
	UfsPath                string
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
