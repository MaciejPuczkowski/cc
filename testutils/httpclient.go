package testutils

import (
	"github.com/MaciejPuczkowski/cc/dict"
	"net/http"
)

type TestRequestData struct {
	Request  *http.Request
	Response *http.Response
}

type TestHTTPClientConf struct {
	URL       string
	Condition func(req *http.Request) bool
	Handle    func(req *http.Request) (*http.Response, error)
	Consume   bool
}

type TestHttpClient struct {
	conf    *dict.Dict[string, TestHTTPClientConf]
	history []TestRequestData
}

func NewTestHttpClient() *TestHttpClient {
	return &TestHttpClient{
		history: []TestRequestData{},
		conf:    dict.New[string, TestHTTPClientConf](),
	}
}

func (c *TestHttpClient) AddConfig(conf TestHTTPClientConf) {

}

func (t *TestHttpClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}
