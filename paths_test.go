package alluxio

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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
	"github.com/Alluxio/alluxio-go/optiontest"
	"github.com/Alluxio/alluxio-go/wire"
	"github.com/Alluxio/alluxio-go/wiretest"
)

func init() {
	seed := time.Now().UnixNano()
	log.Printf("Using seed: %v", seed)
	rand.Seed(seed)
}

// setupClient sets up a instance of the Alluxio file system client connected
// to a local test server that uses the given handler.
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

// jsonHandler creates a handler for the given resource that:
// 1. JSON-decodes the input from the request body
// 2. checks that the received input matches the expected input
// 3. JSON-encodes the given output to the response body
func jsonHandler(t *testing.T, resource string, input, output, expected interface{}) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if got, want := r.URL.Path, join("", apiPrefix, pathsPrefix, resource); got != want {
			t.Fatalf("Unexpected path: got %v, want %v", got, want)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("ReadAll() failed: %v", err)
		}
		decoder := json.NewDecoder(bytes.NewBuffer(body))
		if err := decoder.Decode(&input); err != nil {
			t.Fatalf("Decode() failed: %v\nBody:\n%s", err, body)
		}
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
	input, expected := option.CreateDirectory{}, optiontest.RandomCreateDirectory()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, createDirectory), &input, nil, &expected))
	defer cleanup()
	if err := fs.CreateDirectory(path, &expected); err != nil {
		t.Fatalf("CreateDirectory() failed: %v", err)
	}
}

func TestFileSystemCreateFile(t *testing.T) {
	input, expected := option.CreateFile{}, optiontest.RandomCreateFile()
	output := rand.Int()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, createFile), &input, output, &expected))
	defer cleanup()
	if id, err := fs.CreateFile(path, &expected); err != nil {
		t.Fatalf("CreateFile() failed: %v", err)
	} else if id != output {
		t.Fatalf("Unexpected output: got %v, want %v", id, output)
	}
}

func TestFileSystemDelete(t *testing.T) {
	input, expected := option.Delete{}, optiontest.RandomDelete()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, delete), &input, nil, &expected))
	defer cleanup()
	if err := fs.Delete(path, &expected); err != nil {
		t.Fatalf("Delete() failed: %v", err)
	}
}

func TestFileSystemExists(t *testing.T) {
	input, expected := option.Exists{}, optiontest.RandomExists()
	path := "/foo"
	output := wiretest.RandomBool()
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, exists), &input, output, &expected))
	defer cleanup()
	if ok, err := fs.Exists(path, &expected); err != nil {
		t.Fatalf("Exists() failed: %v", err)
	} else if ok != output {
		t.Fatalf("Unexpected output: got %v, want %v", ok, output)
	}
}

func TestFileSystemFree(t *testing.T) {
	input, expected := option.Free{}, optiontest.RandomFree()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, free), &input, nil, &expected))
	defer cleanup()
	if err := fs.Free(path, &expected); err != nil {
		t.Fatalf("Free() failed: %v", err)
	}
}

func TestFileSystemGetStatus(t *testing.T) {
	input, expected := option.GetStatus{}, optiontest.RandomGetStatus()
	output := wiretest.RandomFileInfo()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, getStatus), &input, output, &expected))
	defer cleanup()
	if fileInfo, err := fs.GetStatus(path, &expected); err != nil {
		t.Fatalf("GetStatus() failed: %v", err)
	} else if !reflect.DeepEqual(&output, fileInfo) {
		t.Fatalf("Unexpected output: got %#v, want %#v", fileInfo, output)
	}
}

func TestFileSystemListStatus(t *testing.T) {
	input, expected := option.ListStatus{}, optiontest.RandomListStatus()
	output := wire.FileInfos(make([]wire.FileInfo, rand.Intn(10)))
	path := "/foo"
	for i := 0; i < len(output); i++ {
		output[i] = wiretest.RandomFileInfo()
	}
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, listStatus), &input, output, &expected))
	defer cleanup()
	if fileInfos, err := fs.ListStatus(path, &expected); err != nil {
		t.Fatalf("ListStatus() failed: %v", err)
	} else if !reflect.DeepEqual(output, fileInfos) {
		t.Fatalf("Unexpected output: got %#v, want %#v", fileInfos, output)
	}
}

func TestFileSystemMount(t *testing.T) {
	input, expected := option.Mount{}, optiontest.RandomMount()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, mount), &input, nil, &expected))
	defer cleanup()
	if err := fs.Mount(path, wiretest.RandomString(), &expected); err != nil {
		t.Fatalf("Mount() failed: %v", err)
	}
}

func TestFileSystemOpenFile(t *testing.T) {
	input, expected := option.OpenFile{}, optiontest.RandomOpenFile()
	output := rand.Int()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, openFile), &input, output, &expected))
	defer cleanup()
	if id, err := fs.OpenFile(path, &expected); err != nil {
		t.Fatalf("OpenFile() failed: %v", err)
	} else if id != output {
		t.Fatalf("Unexpected output: got %v, want %v", id, output)
	}
}

func TestFileSystemRename(t *testing.T) {
	input, expected := option.Rename{}, optiontest.RandomRename()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, rename), &input, nil, &expected))
	defer cleanup()
	if err := fs.Rename(path, wiretest.RandomString(), &expected); err != nil {
		t.Fatalf("Rename() failed: %v", err)
	}
}

func TestFileSystemSetAttribute(t *testing.T) {
	input, expected := option.SetAttribute{}, optiontest.RandomSetAttribute()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, setAttribute), &input, nil, &expected))
	defer cleanup()
	if err := fs.SetAttribute(path, &expected); err != nil {
		t.Fatalf("SetAttribute() failed: %v", err)
	}
}

func TestFileSystemUnmount(t *testing.T) {
	input, expected := option.Unmount{}, optiontest.RandomUnmount()
	path := "/foo"
	fs, cleanup := setupClient(t, jsonHandler(t, join(path, unmount), &input, nil, &expected))
	defer cleanup()
	if err := fs.Unmount(path, &expected); err != nil {
		t.Fatalf("Unmount() failed: %v", err)
	}
}
