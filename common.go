package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultTimeout = 5 * time.Minute
)

const (
	apiPrefix     = "api/v1"
	pathsPrefix   = "paths"
	streamsPrefix = "streams"

	// Path endpoints
	createDirectory = "create-directory"
	createFile      = "create-file"
	delete          = "delete"
	exists          = "exists"
	free            = "free"
	getStatus       = "get-status"
	listStatus      = "list-status"
	mount           = "mount"
	openFile        = "open-file"
	rename          = "rename"
	setAttribute    = "set-attribute"
	unmount         = "unmount"

	// Stream endpoints
	close = "close"
	read  = "read"
	write = "write"
)

type Client struct {
	host   string
	port   int
	prefix string
	http   http.Client
}

func (client *Client) endpointURL(endpoint string, params map[string]string) string {
	paramsStr := []string{}
	for key, value := range params {
		paramsStr = append(paramsStr, key+"="+url.QueryEscape(value))
	}
	return fmt.Sprintf("http://%v:%v/%v/%v?%v", client.host, client.port, client.prefix, endpoint, strings.Join(paramsStr, "&"))
}

func join(components ...string) string {
	return strings.Join(components, "/")
}

func (client *FileSystem) post(method string, params map[string]string, input interface{}, output interface{}) error {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(input); err != nil {
		return fmt.Errorf("Encode() failed: %v", err)
	}
	resp, err := client.http.Post(client.endpointURL(method, params), "application/json", &b)
	if err != nil {
		return err
	}
	if err := check(resp); err != nil {
		return err
	}
	if err := process(resp, output); err != nil {
		return err
	}
	return nil
}

func check(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("%v (%v):\n%v", resp.Status, resp.StatusCode, err)
		}
		return fmt.Errorf("%v (%v):\n%s", resp.Status, resp.StatusCode, bytes)
	}
	return nil
}

func process(resp *http.Response, output interface{}) error {
	defer resp.Body.Close()
	if output != nil {
		contentType := resp.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
				return fmt.Errorf("Decode() failed: %v", err)
			}
		default:
			return fmt.Errorf("Unsupported response type: %v", contentType)
		}
	}
	return nil
}
