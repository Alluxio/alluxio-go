package wiretest

import (
	"math/rand"

	"github.com/alluxio/alluxio-go/wire"
)

// RandomFileInfo generates a random instance of wire.FileInfo.
func RandomFileInfo() wire.FileInfo {
	blockIDs := make([]int64, rand.Intn(10))
	for i := 0; i < len(blockIDs); i++ {
		blockIDs[i] = rand.Int63()
	}
	fileBlockInfos := make([]wire.FileBlockInfo, rand.Intn(10))
	for i := 0; i < len(fileBlockInfos); i++ {
		fileBlockInfos[i] = RandomFileBlockInfo()
	}
	return wire.FileInfo{
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
