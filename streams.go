package alluxio

import (
	"fmt"
	"io"
)

// Close closes a stream.
func (client *Client) Close(id int) error {
	return client.post(join(streamsPrefix, fmt.Sprintf("%d", id), close), nil, nil, nil)
}

// Read reads from a stream.
func (client *Client) Read(id int) (io.ReadCloser, error) {
	suffix := join(streamsPrefix, fmt.Sprintf("%d", id), read)
	resp, err := client.http.Post(client.endpointURL(suffix, nil), "application/json", nil)
	if err != nil {
		return nil, err
	}
	if err := check(resp); err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// Write writes to a stream.
func (client *Client) Write(id int, input io.Reader) (int, error) {
	suffix := join(streamsPrefix, fmt.Sprintf("%d", id), write)
	resp, err := client.http.Post(client.endpointURL(suffix, nil), "application/octet-stream", input)
	if err != nil {
		return -1, err
	}
	if err := check(resp); err != nil {
		return -1, err
	}
	var result int
	if err := process(resp, &result); err != nil {
		return -1, err
	}
	return result, nil
}
