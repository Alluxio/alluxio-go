package alluxio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"testing"
)

var message = []byte("Greetings traveller!")

func streamHandler(t *testing.T, id int, method string, output io.Reader, expected []byte) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.Path, join("", apiPrefix, streamsPrefix, fmt.Sprintf("%d", id), method); got != want {
			t.Fatalf("Unexpected path: got %v, want %v", got, want)
		}
		if expected != nil {
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("ReadAll() failed: %v", err)
			}
			defer r.Body.Close()
			if !reflect.DeepEqual(bytes, expected) {
				t.Fatalf("Unexpected input: got %s, want %s", bytes, expected)
			}
		}
		if output != nil {
			w.Header().Set("Content-Type", "application/octet-stream")
			io.Copy(w, output)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(len(message))
		}
	})
}

func TestFileSystemClose(t *testing.T) {
	id := rand.Int()
	fs, cleanup := setupClient(t, streamHandler(t, id, close, nil, nil))
	defer cleanup()
	if err := fs.Close(id); err != nil {
		t.Fatalf("Close() failed: %v", err)
	}
}

func TestFileSystemRead(t *testing.T) {
	id := rand.Int()
	fs, cleanup := setupClient(t, streamHandler(t, id, read, bytes.NewBuffer(message), nil))
	defer cleanup()
	reader, err := fs.Read(id)
	if err != nil {
		t.Fatalf("Read() failed: %v", err)
	}
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll() failed: %v", err)
	}
	if !reflect.DeepEqual(bytes, message) {
		t.Fatalf("Unexpected output: got %s, want %s", bytes, message)
	}
}

func TestFileSystemWrite(t *testing.T) {
	id := rand.Int()
	fs, cleanup := setupClient(t, streamHandler(t, id, write, nil, message))
	defer cleanup()
	if n, err := fs.Write(id, bytes.NewBuffer(message)); err != nil {
		t.Fatalf("Write() failed: %v", err)
	} else if got, want := n, len(message); got != want {
		t.Fatalf("Unexpected number of bytes written: got %v, want %v", got, want)
	}
}
