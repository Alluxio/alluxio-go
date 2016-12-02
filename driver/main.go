package main

import (
	"fmt"
	"os"
	"time"

	alluxio "github.com/Alluxio/alluxio-go"
)

func main() {
	fs := alluxio.NewFileSystem("localhost", 39999, time.Minute)
	fileInfos, err := fs.ListStatus("/", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	for _, fileInfo := range fileInfos {
		fmt.Printf("%v\n", fileInfo.Name)
	}
}
