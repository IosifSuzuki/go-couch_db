package client

import (
	"net/url"
	"strings"
)

const (
	HttpProtocol  = "http://"
	HttpsProtocol = "https://"
)

type PathBuilder interface {
	AddSegment(string) PathBuilder
	AddQuery(key, value string) PathBuilder
	Build() string
}

type pathBuilder struct {
	protocol     string
	baseURL      string
	pathSegments []string
	query        url.Values
}

func NewPath(protocol, baseURL string) PathBuilder {
	return &pathBuilder{
		protocol:     protocol,
		baseURL:      baseURL,
		pathSegments: make([]string, 0),
		query:        make(url.Values),
	}
}

func (p *pathBuilder) AddSegment(segment string) PathBuilder {
	p.pathSegments = append(p.pathSegments, strings.Trim(segment, "/"))
	return p
}

func (p *pathBuilder) AddQuery(key, value string) PathBuilder {
	p.query.Add(key, value)
	return p
}

func (p *pathBuilder) Build() string {
	pathSegments := strings.Join(append([]string{p.baseURL}, p.pathSegments...), "/")
	fullPath := p.protocol + pathSegments
	if queryParameters := p.query.Encode(); len(queryParameters) > 0 {
		return fullPath + "?" + p.query.Encode()
	}
	return fullPath
}
