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
		isRooted  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "just domain",
			fields: fields{
				port:     80,
				protocol: "http",
				domain:   "example.com",
			},
			want: "http://example.com",
		},
		{
			name: "just domain + suffix",
			fields: fields{
				hasSuffix: true,
				port:      80,
				protocol:  "http",
				domain:    "example.com",
			},
			want: "http://example.com/",
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
			want: "http://example.com/?baz=qux",
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
			want: "http://example.com/?baz=qux#foo",
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
			want: "http://example.com/foo/bar/?baz=qux#foo",
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
			want: "http://example.com/foo/bar?baz=qux#foo",
		},
		{
			name: "just domain + query + fragment + path + rooted",
			fields: fields{
				isRooted:  true,
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
			want: "http://example.com/foo/bar?baz=qux#foo",
		},
		{
			name: "just  path",
			fields: fields{
				hasSuffix: false,
				port:      80,
				path:      []string{"foo", "bar"},
			},
			want: "foo/bar",
		},
		{
			name: "just  path",
			fields: fields{
				hasSuffix: false,
				port:      80,
				path:      []string{"foo", "bar"},
				isRooted:  true,
			},
			want: "/foo/bar",
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
				isRooted:  tt.fields.isRooted,
			}
			got := b.URL()
			if got != tt.want {
				t.Errorf("URLBuilder.URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
