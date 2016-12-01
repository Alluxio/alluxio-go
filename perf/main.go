//+build ignore

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
	GB     = 1024 * 1024 * 1024
	numOps = 100
)

func main() {
	switch os.Args[1] {
	case "metadata":
		metadataBenchmark()
	case "read":
		readBenchmark()
	case "write":
		writeBenchmark()
	}
}

func deleteIfExists(fs *alluxio.FileSystem, path string) {
	if ok, err := fs.Exists(path); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	} else if ok {
		if err := fs.Delete(path); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
	}
}

func metadataBenchmark() {
	fs := alluxio.NewFileSystem("localhost", 39999, time.Minute)
	start := time.Now()
	for i := 0; i < numOps; i++ {
		if _, err := fs.Exists("/data"); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
	}
	end := time.Now()
	fmt.Printf("Performed %v operations in %v seconds\n", numOps, end.Sub(start).Seconds())
}

func readBenchmark() {
	fs := alluxio.NewFileSystem("localhost", 39999, time.Minute)
	deleteIfExists(fs, "/data")
	start := time.Now()
	buffer := make([]byte, GB)
	if err := fs.Upload("/data", nopCloser{bytes.NewBuffer(buffer)}); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	for i := 0; i < numOps; i++ {
		if reader, err := fs.Download("/data"); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		} else {
			if _, err := io.Copy(ioutil.Discard, reader); err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				os.Exit(-1)
			}
		}
	}
	end := time.Now()
	fmt.Printf("Read %vGBs in %v seconds\n", numOps, end.Sub(start).Seconds())
}

func writeBenchmark() {
	fs := alluxio.NewFileSystem("localhost", 39999, time.Minute)
	buffer := make([]byte, GB)
	for i := 0; i < numOps; i++ {
		deleteIfExists(fs, fmt.Sprintf("/data-%d", i+1))
	}
	start := time.Now()
	for i := 0; i < numOps; i++ {
		if err := fs.Upload(fmt.Sprintf("/data-%d", i+1), nopCloser{bytes.NewBuffer(buffer)}); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
	}
	end := time.Now()
	fmt.Printf("Wrote %vGBs in %v seconds\n", numOps, end.Sub(start).Seconds())
}
