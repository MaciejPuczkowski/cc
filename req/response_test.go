package req

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestResponseReader_Read(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		rr      *ResponseReader
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Simple test",
			rr:   NewResponseReader(),
			args: args{
				r: &http.Response{
					StatusCode: 200,
					Body:       http.NoBody,
					Header:     http.Header{},
				},
			},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "Simple test with content",
			rr:   NewResponseReader(),
			args: args{
				r: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader("Hello, world!")),
					Header:     http.Header{},
				},
			},
			want:    []byte("Hello, world!"),
			wantErr: false,
		},
		{
			name: "Simple test with error and content",
			rr:   NewResponseReader(),
			args: args{
				r: &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(strings.NewReader("Hello, world!")),
					Header:     http.Header{},
				},
			},
			want:    []byte("Hello, world!"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &ResponseReader{}
			got, err := rr.Read(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResponseReader.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResponseReader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
