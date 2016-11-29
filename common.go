package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultTimeout = 5 * time.Minute
)

const (
	fileSystemServicePrefix = "v1/api/file"

	// Common endpoints
	serviceName    = "service_name"
	serviceVersion = "service_version"

	// File System endpoints
	createDirectory = "create_directory"
	delete          = "delete"
	download        = "download"
	exists          = "exists"
	free            = "free"
	getStatus       = "status"
	listStatus      = "list_status"
	mount           = "mount"
	rename          = "rename"
	setAttribute    = "set_attribute"
	unmount         = "unmount"
	upload          = "upload"
)

type Client struct {
	host   string
	port   int
	prefix string
	http   http.Client
}

func (client *Client) endpointURL(endpoint string, params map[string]string) string {
	paramsStr := ""
	for key, value := range params {
		paramsStr += key + "=" + url.QueryEscape(value) + "&"
	}
	return "http://" + client.host + ":" + fmt.Sprintf("%v", client.port) + "/" + client.prefix + "/" + endpoint + "?" + paramsStr
}

func (client *FileSystem) get(method string, params map[string]string, result interface{}) error {
	resp, err := client.http.Get(client.endpointURL(method, params))
	if err != nil {
		return err
	}
	if err := checkResponse(resp); err != nil {
		return err
	}
	if err := processResponse(resp, result); err != nil {
		return err
	}
	return nil
}

func (client *FileSystem) post(method string, params map[string]string) error {
	resp, err := client.http.Post(client.endpointURL(method, params), "application/json", nil)
	if err != nil {
		return err
	}
	if err := checkResponse(resp); err != nil {
		return err
	}
	if err := processResponse(resp, nil); err != nil {
		return err
	}
	return nil
}

func (client *Client) ServiceName() (string, error) {
	resp, err := http.Get(client.endpointURL(serviceName, nil))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result string
	if err := processResponse(resp, &result); err != nil {
		return "", err
	}
	return result, nil
}

func (client *Client) ServiceVersion() (int64, error) {
	resp, err := http.Get(client.endpointURL(serviceVersion, nil))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	var result int64
	if err := processResponse(resp, &result); err != nil {
		return -1, err
	}
	return result, nil
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("%v", resp.Status)
		}
		return fmt.Errorf("%s", bytes)
	}
	return nil
}

func processResponse(resp *http.Response, data interface{}) error {
	defer resp.Body.Close()
	if data != nil {
		contentType := resp.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
				return fmt.Errorf("Decode() failed with: %v", err)
			}
		default:
			return fmt.Errorf("Unsupported response type: %v", contentType)
		}
	}
	return nil
}
