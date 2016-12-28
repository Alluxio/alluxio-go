package alluxio

import (
	"github.com/TachyonNexus/alluxio-go/option"
	"github.com/TachyonNexus/alluxio-go/wire"
)

var (
	noParams = map[string]string{}
)

// CreateDirectory creates a directory.
func (client *Client) CreateDirectory(path string, options *option.CreateDirectory) error {
	return client.post(join(pathsPrefix, path, createDirectory), noParams, options, nil)
}

// CreateFile creates a file, returning a stream id.
func (client *Client) CreateFile(path string, options *option.CreateFile) (int, error) {
	var result int
	if err := client.post(join(pathsPrefix, path, createFile), noParams, options, &result); err != nil {
		return -1, err
	}
	return result, nil
}

// Delete deletes a path.
func (client *Client) Delete(path string, options *option.Delete) error {
	return client.post(join(pathsPrefix, path, delete), noParams, options, nil)
}

// Exists checks whether a path exists.
func (client *Client) Exists(path string, options *option.Exists) (bool, error) {
	var result bool
	if err := client.post(join(pathsPrefix, path, exists), noParams, options, &result); err != nil {
		return false, err
	}
	return result, nil
}

// Free frees a path.
func (client *Client) Free(path string, options *option.Free) error {
	return client.post(join(pathsPrefix, path, free), noParams, options, nil)
}

// GetStatus gets a path status.
func (client *Client) GetStatus(path string, options *option.GetStatus) (*wire.FileInfo, error) {
	var result wire.FileInfo
	if err := client.post(join(pathsPrefix, path, getStatus), noParams, options, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListStatus lists path statuses.
func (client *Client) ListStatus(path string, options *option.ListStatus) (wire.FileInfos, error) {
	var result wire.FileInfos
	if err := client.post(join(pathsPrefix, path, listStatus), noParams, options, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Mount mounts a UFS source to a path.
func (client *Client) Mount(path, src string, options *option.Mount) error {
	params := map[string]string{
		"src": src,
	}
	return client.post(join(pathsPrefix, path, mount), params, options, nil)
}

// OpenFile opens a file, returning a stream id.
func (client *Client) OpenFile(path string, options *option.OpenFile) (int, error) {
	var result int
	if err := client.post(join(pathsPrefix, path, openFile), noParams, options, &result); err != nil {
		return -1, err
	}
	return result, nil
}

// Rename renames a path.
func (client *Client) Rename(path, dst string, options *option.Rename) error {
	params := map[string]string{
		"dst": dst,
	}
	return client.post(join(pathsPrefix, path, rename), params, options, nil)
}

// SetAttribute sets path attributes.
func (client *Client) SetAttribute(path string, options *option.SetAttribute) error {
	return client.post(join(pathsPrefix, path, setAttribute), noParams, options, nil)
}

// Unmount unmounts a path.
func (client *Client) Unmount(path string, options *option.Unmount) error {
	return client.post(join(pathsPrefix, path, unmount), noParams, options, nil)
}
