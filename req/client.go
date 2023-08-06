package req

import (
	"bytes"
	"net/http"
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client completes http.Client with missing methods and unpacks the response
type Client struct {
	http           HTTPDoer
	responseReader *ResponseReader
}

// NewClient creates a new client
func NewClient(http HTTPDoer) *Client {
	return &Client{
		http:           http,
		responseReader: NewResponseReader(),
	}
}
func NewClientDefault() *Client {
	return NewClient(http.DefaultClient)
}

// Get performs a GET request
func (c *Client) Get(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.send(request)
}

// Post performs a POST request
func (c *Client) Post(url string, body []byte) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return c.send(request)
}

// Put performs a PUT request
func (c *Client) Put(url string, body []byte) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return c.send(request)
}

// Delete performs a DELETE request
func (c *Client) Delete(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return c.send(request)
}

// Patch performs a PATCH request
func (c *Client) Patch(url string, body []byte) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return c.send(request)
}

func (c *Client) send(request *http.Request) ([]byte, error) {
	response, err := c.http.Do(request)
	if err != nil {
		return nil, err
	}
	return c.responseReader.Read(response)
}
