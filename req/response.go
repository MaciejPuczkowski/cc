package req

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	ErrStatusNotOK = errors.New("unexpected status code")
)

type ResponseReader struct{}

func NewResponseReader() *ResponseReader {
	return &ResponseReader{}
}

func (rr *ResponseReader) Read(r *http.Response) ([]byte, error) {
	var message []byte
	var err error
	message, err = io.ReadAll(r.Body)
	if err != nil {
		message = []byte{}
	}
	if r.StatusCode >= 400 {
		return message, ErrStatusNotOK
	}
	return message, nil
}

// JSONResponseReader is a ResponseReader that unmarshals the response body into a JSON object.
type JSONResponseReader[T any] struct {
	rr *ResponseReader
}

// NewJSONResponseReader returns a new JSONResponseReader.
func NewJSONResponseReader[T any]() *JSONResponseReader[T] {
	return &JSONResponseReader[T]{
		rr: NewResponseReader(),
	}
}

// Read unmarshals the response body into v.
func (jrr *JSONResponseReader[T]) Read(r *http.Response, v *T) error {
	message, err := jrr.rr.Read(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(message, v)
}

type TextResponseReader struct {
	rr *ResponseReader
}

func NewTextResponseReader() *TextResponseReader {
	return &TextResponseReader{
		rr: NewResponseReader(),
	}
}

// Read unmarshals the response body into v.
func (trr *TextResponseReader) Read(r *http.Response, v *string) error {
	message, err := trr.rr.Read(r)
	if err != nil {
		return err
	}
	*v = string(message)
	return nil
}
