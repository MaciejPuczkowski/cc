package testutils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// TestRequestData is a struct holding Request and Response
type TestRequestData struct {
	Request  *http.Request
	Response *http.Response
	Err      error
}

type TestHTTPClientConf struct {
	URL          string
	Method       string
	Condition    func(req *http.Request) bool
	HandlerFunc  http.HandlerFunc
	Handler      http.Handler
	ConsumeAfter int
}

type TestHttpClient struct {
	conf    []TestHTTPClientConf
	history []TestRequestData
}

func NewTestHttpClient() *TestHttpClient {
	return &TestHttpClient{
		history: []TestRequestData{},
		conf:    make([]TestHTTPClientConf, 0),
	}
}

func (c *TestHttpClient) AddConfig(conf TestHTTPClientConf) {
	c.conf = append(c.conf, conf)
}

func (t *TestHttpClient) Do(req *http.Request) (*http.Response, error) {
	var result TestRequestData
	result.Request = req
	for i, conf := range t.conf {
		if t.matches(conf, req) {
			result.Response, result.Err = t.handle(&conf, req)
			if conf.ConsumeAfter == 1 {
				t.conf = append(t.conf[:i], t.conf[i+1:]...)
			}
			if conf.ConsumeAfter > 1 {
				conf.ConsumeAfter--
			}
			t.history = append(t.history, result)
			return result.Response, result.Err
		}
	}
	return &http.Response{}, fmt.Errorf("%w: %s", ErrRouteNotFound, req.URL.String())
}
func (t *TestHttpClient) matches(conf TestHTTPClientConf, req *http.Request) bool {
	if conf.Method != "" && conf.Method != req.Method {
		return false
	}
	if conf.URL != "" && !t.matchesURL(conf, req) {
		return false
	}
	if conf.Condition != nil && !conf.Condition(req) {
		return false
	}
	return true
}

func (t *TestHttpClient) matchesURL(conf TestHTTPClientConf, req *http.Request) bool {
	matches, err := regexp.MatchString(conf.URL, req.URL.String())
	return err == nil && matches
}

func (t *TestHttpClient) handle(conf *TestHTTPClientConf, req *http.Request) (*http.Response, error) {
	rw := newResponseWriter()
	if conf.HandlerFunc != nil {
		conf.HandlerFunc(rw, req)
	} else if conf.Handler != nil {
		conf.Handler.ServeHTTP(rw, req)
	}
	var response response
	rw.ScanResponse(&response)
	return &http.Response{
		StatusCode: response.StatusCode,
		Body:       io.NopCloser(bytes.NewReader(response.Body)),
		Header:     response.Header,
	}, nil
}

func (t *TestHttpClient) History() []TestRequestData {
	return t.history
}

func (t *TestHttpClient) LastRequestResult() *TestRequestData {
	if t.history == nil {
		return nil
	}
	return &t.history[len(t.history)-1]
}
