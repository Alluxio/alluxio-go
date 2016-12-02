package client

import (
	"encoding/json"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/Alluxio/alluxio-go/option"
	"github.com/Alluxio/alluxio-go/wire"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func setupClient(t *testing.T, handler http.Handler) (*Client, func()) {
	testServer := httptest.NewServer(handler)
	url, err := url.Parse(testServer.URL)
	if err != nil {
		t.Fatalf("Parse(%v) failed: %v", testServer.URL, err)
	}
	host, portStr, err := net.SplitHostPort(url.Host)
	if err != nil {
		t.Fatalf("SplitHostPort(%v) failed: %v", url.Host, err)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		t.Fatalf("Atoi(%v) failed: %v", portStr, err)
	}
	return NewClient(host, port, time.Minute), func() {
		testServer.Close()
	}
}

func jsonHandler(t *testing.T, path, method string, input, output, expected interface{}) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, join("", apiPrefix, pathsPrefix, path, method); got != want {
			t.Fatalf("Unexpected path: got %v, want %v", got, want)
		}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&input); err != nil {
			t.Fatalf("Decode() failed: %v", err)
		}
		defer r.Body.Close()
		if !reflect.DeepEqual(input, expected) {
			t.Fatalf("Unexpected input: got %v, want %v", input, expected)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(output); err != nil {
			t.Fatalf("Encode() failed: %v", err)
		}
	})
}

func TestFileSystemCreateDirectory(t *testing.T) {
	input, expected := option.CreateDirectory{}, option.RandomCreateDirectory()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, createDirectory, &input, nil, &expected))
	defer cleanup()
	if err := fs.CreateDirectory(path, &expected); err != nil {
		t.Fatalf("CreateDirectory() failed: %v", err)
	}
}

func TestFileSystemCreateFile(t *testing.T) {
	input, expected := option.CreateFile{}, option.RandomCreateFile()
	output := rand.Int()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, createFile, &input, output, &expected))
	defer cleanup()
	if id, err := fs.CreateFile(path, &expected); err != nil {
		t.Fatalf("CreateFile() failed: %v", err)
	} else if id != output {
		t.Fatalf("Unexpected output: got %v, want %v", id, output)
	}
}

func TestFileSystemDelete(t *testing.T) {
	input, expected := option.Delete{}, option.RandomDelete()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, delete, &input, nil, &expected))
	defer cleanup()
	if err := fs.Delete(path, &expected); err != nil {
		t.Fatalf("Delete() failed: %v", err)
	}
}

func TestFileSystemExists(t *testing.T) {
	input, expected := option.Exists{}, option.RandomExists()
	path := "/foo"
	output := wire.RandomBool()
	fs, cleanup := setupClient(t, jsonHandler(t, path, exists, &input, output, &expected))
	defer cleanup()
	if ok, err := fs.Exists(path, &expected); err != nil {
		t.Fatalf("Exists() failed: %v", err)
	} else if ok != output {
		t.Fatalf("Unexpected output: got %v, want %v", ok, output)
	}
}

func TestFileSystemFree(t *testing.T) {
	input, expected := option.Free{}, option.RandomFree()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, free, &input, nil, &expected))
	defer cleanup()
	if err := fs.Free(path, &expected); err != nil {
		t.Fatalf("Free() failed: %v", err)
	}
}

func TestFileSystemGetStatus(t *testing.T) {
	input, expected := option.GetStatus{}, option.RandomGetStatus()
	output := wire.RandomFileInfo()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, getStatus, &input, output, &expected))
	defer cleanup()
	if fileInfo, err := fs.GetStatus(path, &expected); err != nil {
		t.Fatalf("GetStatus() failed: %v", err)
	} else if !reflect.DeepEqual(&output, fileInfo) {
		t.Fatalf("Unexpected output: got %#v, want %#v", fileInfo, output)
	}
}

func TestFileSystemListStatus(t *testing.T) {
	input, expected := option.ListStatus{}, option.RandomListStatus()
	output := wire.FileInfos(make([]wire.FileInfo, rand.Intn(10)))
	path := "/foo"
	for i := 0; i < len(output); i++ {
		output[i] = wire.RandomFileInfo()
	}
	fs, cleanup := setupClient(t, jsonHandler(t, path, listStatus, &input, output, &expected))
	defer cleanup()
	if fileInfos, err := fs.ListStatus(path, &expected); err != nil {
		t.Fatalf("ListStatus() failed: %v", err)
	} else if !reflect.DeepEqual(output, fileInfos) {
		t.Fatalf("Unexpected output: got %#v, want %#v", fileInfos, output)
	}
}

func TestFileSystemMount(t *testing.T) {
	input, expected := option.Mount{}, option.RandomMount()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, mount, &input, nil, &expected))
	defer cleanup()
	if err := fs.Mount(path, wire.RandomString(), &expected); err != nil {
		t.Fatalf("Mount() failed: %v", err)
	}
}

func TestFileSystemOpenFile(t *testing.T) {
	input, expected := option.OpenFile{}, option.RandomOpenFile()
	output := rand.Int()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, openFile, &input, output, &expected))
	defer cleanup()
	if id, err := fs.OpenFile(path, &expected); err != nil {
		t.Fatalf("OpenFile() failed: %v", err)
	} else if id != output {
		t.Fatalf("Unexpected output: got %v, want %v", id, output)
	}
}

func TestFileSystemRename(t *testing.T) {
	input, expected := option.Rename{}, option.RandomRename()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, rename, &input, nil, &expected))
	defer cleanup()
	if err := fs.Rename(path, wire.RandomString(), &expected); err != nil {
		t.Fatalf("Rename() failed: %v", err)
	}
}

func TestFileSystemSetAttribute(t *testing.T) {
	input, expected := option.SetAttribute{}, option.RandomSetAttribute()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, setAttribute, &input, nil, &expected))
	defer cleanup()
	if err := fs.SetAttribute(path, &expected); err != nil {
		t.Fatalf("SetAttribute() failed: %v", err)
	}
}

func TestFileSystemUnmount(t *testing.T) {
	input, expected := option.Unmount{}, option.RandomUnmount()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, path, unmount, &input, nil, &expected))
	defer cleanup()
	if err := fs.Unmount(path, &expected); err != nil {
		t.Fatalf("Unmount() failed: %v", err)
	}
}
