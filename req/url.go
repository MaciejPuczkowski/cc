package req

import (
	"fmt"
	"strings"
)

type URLBuilder struct {
	protocol  string
	domain    string
	port      int
	path      []string
	fragment  string
	query     map[string][]string
	hasSuffix bool
}

func NewURLBuilder() *URLBuilder {
	return &URLBuilder{
		protocol: "http",
		domain:   "localhost",
		port:     80,
		path:     []string{""},
		fragment: "",
		query:    make(map[string][]string),
	}
}

func (b *URLBuilder) SetSuffix(has bool) *URLBuilder {
	b.hasSuffix = has
	return b
}

func (b *URLBuilder) SetProtocol(protocol string) *URLBuilder {
	b.protocol = protocol
	return b
}

func (b *URLBuilder) SetDomain(domain string) *URLBuilder {
	b.domain = domain
	return b
}

func (b *URLBuilder) SetPort(port int) *URLBuilder {
	b.port = port
	return b
}

func (b *URLBuilder) SetPath(path string) *URLBuilder {
	path = strings.Trim(path, "/")
	b.path = strings.Split(path, "/")
	return b
}

func (b *URLBuilder) SetFragment(fragment string) *URLBuilder {
	b.fragment = fragment
	return b
}

func (b *URLBuilder) AddQuery(key string, value string) *URLBuilder {
	if _, ok := b.query[key]; !ok {
		b.query[key] = []string{}
	}
	b.query[key] = append(b.query[key], value)
	return b
}

func (b *URLBuilder) AppendPath(path string) *URLBuilder {
	path = strings.Trim(path, "/")
	b.path = append(b.path, strings.Split(path, "/")...)
	return b
}

func (b *URLBuilder) Copy() *URLBuilder {
	return &URLBuilder{
		protocol: b.protocol,
		domain:   b.domain,
		port:     b.port,
		path:     b.path,
		fragment: b.fragment,
		query:    b.query,
	}
}

func (b *URLBuilder) Build() (string, error) {
	domain := strings.Trim(b.domain, "/")
	protocol := strings.Trim(b.protocol, "/")
	if domain == "" {
		return "", fmt.Errorf("%w: domain is empty", ErrURLBuild)
	}
	if protocol == "" {
		return "", fmt.Errorf("%w: protocol is empty", ErrURLBuild)
	}
	url := protocol + "://" + domain
	if b.port != 80 {
		url = fmt.Sprintf("%s:%d", url, b.port)
	}
	if len(b.path) > 0 {
		url += "/" + strings.Join(b.path, "/")
	}
	if b.hasSuffix {
		url += "/"
	}
	if len(b.query) > 0 {
		url += "?"
		for key, values := range b.query {
			for _, value := range values {
				url += key + "=" + value + "&"
			}
		}
		url = strings.TrimRight(url, "&")
	}
	if b.fragment != "" {
		url += "#" + b.fragment
	}
	return url, nil
}
