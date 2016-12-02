package client

import (
	"net/http"
	"time"
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
		prefix: apiPrefix,
		http: http.Client{
			Timeout: timeout,
		},
	}}
}
