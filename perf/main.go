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
	if ok, err := fs.Exists(path, nil); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	} else if ok {
		if err := fs.Delete(path, nil); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
	}
}

func readFile(fs *alluxio.FileSystem, path string) {
	id, err := fs.OpenFile(path, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	reader, err := fs.Read(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	if _, err := io.Copy(ioutil.Discard, reader); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	if err := fs.Close(id); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func writeFile(fs *alluxio.FileSystem, path string) {
	id, err := fs.CreateFile(path, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	buffer := make([]byte, GB)
	if _, err := fs.Write(id, nopCloser{bytes.NewBuffer(buffer)}); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	if err := fs.Close(id); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func metadataBenchmark() {
	fs := alluxio.NewFileSystem("localhost", 39999, time.Minute)
	start := time.Now()
	for i := 0; i < numOps; i++ {
		if _, err := fs.Exists("/data", nil); err != nil {
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
	writeFile(fs, "/data")
	start := time.Now()
	for i := 0; i < numOps; i++ {
		readFile(fs, "/data")
	}
	end := time.Now()
	fmt.Printf("Read %vGBs in %v seconds\n", numOps, end.Sub(start).Seconds())
}

func writeBenchmark() {
	fs := alluxio.NewFileSystem("localhost", 39999, time.Minute)
	for i := 0; i < numOps; i++ {
		deleteIfExists(fs, fmt.Sprintf("/data-%d", i+1))
	}
	start := time.Now()
	for i := 0; i < numOps; i++ {
		writeFile(fs, fmt.Sprintf("/data-%d", i+1))
	}
	end := time.Now()
	fmt.Printf("Wrote %vGBs in %v seconds\n", numOps, end.Sub(start).Seconds())
}
