package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	alluxio "github.com/Alluxio/alluxio-go"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

const (
	GB       = 1024 * 1024 * 1024
	numFiles = 100
)

func main() {
	fs := alluxio.NewFileSystem("localhost", 39999, time.Minute)
	buffer := make([]byte, GB)
	for i := 0; i < numFiles; i++ {
		if err := fs.Delete(fmt.Sprintf("/data-%d", i+1)); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
	}
	start := time.Now()
	for i := 0; i < numFiles; i++ {
		if err := fs.Upload(fmt.Sprintf("/data-%d", i+1), nopCloser{bytes.NewBuffer(buffer)}); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
	}
	end := time.Now()
	fmt.Printf("Wrote %vGBs in %v seconds\n", numFiles, end.Sub(start).Seconds())
}
