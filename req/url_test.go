package req

import "testing"

func TestURLBuilder_Build(t *testing.T) {
	type fields struct {
		protocol  string
		domain    string
		port      int
		path      []string
		fragment  string
		query     map[string][]string
		hasSuffix bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "just domain",
			fields: fields{
				port:     80,
				protocol: "http",
				domain:   "example.com",
			},
			want:    "http://example.com",
			wantErr: false,
		},
		{
			name: "just domain + suffix",
			fields: fields{
				hasSuffix: true,
				port:      80,
				protocol:  "http",
				domain:    "example.com",
			},
			want:    "http://example.com/",
			wantErr: false,
		},
		{
			name: "just domain + suffix + query",
			fields: fields{
				hasSuffix: true,
				port:      80,
				protocol:  "http",
				domain:    "example.com",
				query: map[string][]string{
					"baz": {"qux"},
				},
			},
			want:    "http://example.com/?baz=qux",
			wantErr: false,
		},
		{
			name: "just domain + suffix + query + fragment",
			fields: fields{
				hasSuffix: true,
				port:      80,
				protocol:  "http",
				domain:    "example.com",
				query: map[string][]string{
					"baz": {"qux"},
				},
				fragment: "foo",
			},
			want:    "http://example.com/?baz=qux#foo",
			wantErr: false,
		},
		{
			name: "just domain + suffix + query + fragment + path",
			fields: fields{
				hasSuffix: true,
				port:      80,
				protocol:  "http",
				domain:    "example.com",
				query: map[string][]string{
					"baz": {"qux"},
				},
				path:     []string{"foo", "bar"},
				fragment: "foo",
			},
			want:    "http://example.com/foo/bar/?baz=qux#foo",
			wantErr: false,
		},
		{
			name: "just domain + query + fragment + path",
			fields: fields{
				hasSuffix: false,
				port:      80,
				protocol:  "http",
				domain:    "example.com",
				query: map[string][]string{
					"baz": {"qux"},
				},
				path:     []string{"foo", "bar"},
				fragment: "foo",
			},
			want:    "http://example.com/foo/bar?baz=qux#foo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &URLBuilder{
				protocol:  tt.fields.protocol,
				domain:    tt.fields.domain,
				port:      tt.fields.port,
				path:      tt.fields.path,
				fragment:  tt.fields.fragment,
				query:     tt.fields.query,
				hasSuffix: tt.fields.hasSuffix,
			}
			got, err := b.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("URLBuilder.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
