package alluxio

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/Alluxio/alluxio-go/option"
)

func ExampleClient_download() {
	fs := NewClient("localhost", 39999, time.Minute)
	id, err := fs.OpenFile("/test_file", &option.OpenFile{})
	if err != nil {
		log.Fatal(err)
	}
	r, err := fs.Read(id)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	defer fs.Close(id)
	if _, err := io.Copy(ioutil.Discard, r); err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_exists() {
	fs := NewClient("localhost", 39999, time.Minute)
	ok, err := fs.Exists("/test_path", &option.Exists{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}

func ExampleClient_listStatus() {
	fs := NewClient("localhost", 39999, time.Minute)
	fileInfos, err := fs.ListStatus("/test_path", &option.ListStatus{})
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name)
	}
}

func ExampleClient_upload() {
	fs := NewClient("localhost", 39999, time.Minute)
	id, err := fs.CreateFile("/test_file", &option.CreateFile{})
	if err != nil {
		log.Fatal(err)
	}
	n, err := fs.Write(id, bytes.NewBuffer(make([]byte, 1024)))
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close(id)
	fmt.Println(n)
}
