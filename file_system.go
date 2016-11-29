package client

import (
	"io"
	"net/http"
	"time"

	"github.com/Alluxio/alluxio-go/wire"
)

// FileSystem represents the file system client.
type FileSystem struct {
	Client
}

// NewFileSystem is the file system client factory.
func NewFileSystem(host string, port int, timeout time.Duration) *FileSystem {
	return &FileSystem{Client{
		host:   host,
		port:   port,
		prefix: fileSystemServicePrefix,
		http: http.Client{
			Timeout: timeout,
		},
	}}
}

func (client *FileSystem) CreateDirectory(path string, options ...func(map[string]string)) error {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	return client.post(createDirectory, params)
}

func (client *FileSystem) Delete(path string, options ...func(map[string]string)) error {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	return client.post(delete, params)
}

func (client *FileSystem) Download(path string, options ...func(map[string]string)) (io.ReadCloser, error) {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	resp, err := client.http.Get(client.endpointURL(download, params))
	if err != nil {
		return nil, err
	}
	if err := checkResponse(resp); err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (client *FileSystem) Exists(path string, options ...func(map[string]string)) (bool, error) {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	var result bool
	if err := client.get(exists, params, &result); err != nil {
		return false, err
	}
	return result, nil
}

func (client *FileSystem) Free(path string, options ...func(map[string]string)) error {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	return client.post(free, params)
}

func (client *FileSystem) GetStatus(path string, options ...func(map[string]string)) (*wire.FileInfo, error) {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	var result wire.FileInfo
	if err := client.get(getStatus, params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (client *FileSystem) ListStatus(path string, options ...func(map[string]string)) (wire.FileInfos, error) {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	var result wire.FileInfos
	if err := client.get(listStatus, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (client *FileSystem) Mount(path, ufsPath string, options ...func(map[string]string)) error {
	params := map[string]string{
		"path":    path,
		"ufsPath": ufsPath,
	}
	for _, option := range options {
		option(params)
	}
	return client.post(mount, params)
}

func (client *FileSystem) Rename(srcPath, dstPath string, options ...func(map[string]string)) error {
	params := map[string]string{
		"srcPath": srcPath,
		"dstPath": dstPath,
	}
	for _, option := range options {
		option(params)
	}
	return client.post(rename, params)
}

func (client *FileSystem) SetAttribute(path string, options ...func(map[string]string)) error {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	return client.post(setAttribute, params)
}

func (client *FileSystem) Unmount(path string, options ...func(map[string]string)) error {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	return client.post(unmount, params)
}

func (client *FileSystem) Upload(path string, data io.ReadCloser, options ...func(map[string]string)) error {
	params := map[string]string{
		"path": path,
	}
	for _, option := range options {
		option(params)
	}
	resp, err := client.http.Post(client.endpointURL(upload, params), "application/octet-stream", data)
	if err != nil {
		return err
	}
	if err := checkResponse(resp); err != nil {
		return err
	}
	if err := processResponse(resp, nil); err != nil {
		return err
	}
	return nil
}
