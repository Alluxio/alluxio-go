package wire

import (
	"math/rand"
	"sort"
)

// FileInfo represents a file information.
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

// RandomFileInfo generates a random file information for testing.
func RandomFileInfo() FileInfo {
	blockIDs := make([]int64, rand.Intn(10))
	for i := 0; i < len(blockIDs); i++ {
		blockIDs[i] = rand.Int63()
	}
	fileBlockInfos := make([]FileBlockInfo, rand.Intn(10))
	for i := 0; i < len(fileBlockInfos); i++ {
		fileBlockInfos[i] = RandomFileBlockInfo()
	}
	return FileInfo{
		BlockIds:               blockIDs,
		BlockSizeBytes:         rand.Int63(),
		Cacheable:              RandomBool(),
		Completed:              RandomBool(),
		CreationTimeMs:         rand.Int63(),
		FileBlockInfos:         fileBlockInfos,
		FileID:                 rand.Int63(),
		Folder:                 RandomBool(),
		Group:                  RandomString(),
		InMemoryPercentage:     rand.Int31(),
		LastModificationTimeMs: rand.Int63(),
		Length:                 rand.Int63(),
		Mode:                   rand.Int31(),
		MountPoint:             RandomBool(),
		Name:                   RandomString(),
		Owner:                  RandomString(),
		Path:                   RandomString(),
		Persisted:              RandomBool(),
		PersistenceState:       RandomString(),
		Pinned:                 RandomBool(),
		TTL:                    rand.Int63(),
		TTLAction:              RandomString(),
		UfsPath:                RandomString(),
	}
}
