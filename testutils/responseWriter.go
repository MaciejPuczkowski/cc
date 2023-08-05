package testutils

import "net/http"

type response struct {
	StatusCode int
	Body       []byte
	Header     http.Header
}
type responseWriter struct {
	response response
}

func newResponseWriter() *responseWriter {
	return &responseWriter{
		response: response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       make([]byte, 0),
		},
	}
}

func (rw *responseWriter) Header() http.Header {
	return rw.response.Header
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.response.Body = append(rw.response.Body, b...)
	return len(b), nil
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.response.StatusCode = statusCode
}

func (rw *responseWriter) ScanResponse(r *response) {
	r.Body = rw.response.Body
	r.Header = rw.response.Header
	r.StatusCode = rw.response.StatusCode
}
