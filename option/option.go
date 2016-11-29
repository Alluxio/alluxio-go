package client

import "fmt"

func AllowExists(allowExists bool) func(map[string]string) {
	return set("allowExists", allowExists)
}

func BlockSize(blockSize int64) func(map[string]string) {
	return set("blockSize", blockSize)
}

func Group(group string) func(map[string]string) {
	return set("group", group)
}

func LoadMetadataType(loadMetadataType string) func(map[string]string) {
	return set("loadMetadataType", loadMetadataType)
}

func LocationPolicy(locationPolicy string) func(map[string]string) {
	return set("locationPolicy", locationPolicy)
}

func Mode(mode int16) func(map[string]string) {
	return set("mode", mode)
}

func Owner(owner string) func(map[string]string) {
	return set("owner", owner)
}

func Persisted(persisted bool) func(map[string]string) {
	return set("persisted", persisted)
}

func Pinned(pinned bool) func(map[string]string) {
	return set("pinned", pinned)
}

func ReadOnly(readOnly bool) func(map[string]string) {
	return set("readOnly", readOnly)
}

func ReadType(readType string) func(map[string]string) {
	return set("readType", readType)
}

func Recursive(recursive bool) func(map[string]string) {
	return set("recursive", recursive)
}

func Shared(shared bool) func(map[string]string) {
	return set("shared", shared)
}

func TTL(ttl int64) func(map[string]string) {
	return set("ttl", ttl)
}

func TTLAction(ttlAction string) func(map[string]string) {
	return set("ttlAction", ttlAction)
}

func WriteType(writeType string) func(map[string]string) {
	return set("writeType", writeType)
}

func set(field string, value interface{}) func(map[string]string) {
	return func(params map[string]string) {
		params[field] = fmt.Sprintf("%v", value)
	}
}
