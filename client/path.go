package client

import (
	"github.com/IosifSuzuki/go-couch_db/model"
	"github.com/IosifSuzuki/go-couch_db/tool"
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
	AddQueryModel(params model.QueryViewParams) PathBuilder
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

func (p *pathBuilder) AddQueryModel(params model.QueryViewParams) PathBuilder {
	urlValues, err := tool.BuildUrlValues(params)
	if err != nil {
		panic(err)
	}
	for key, values := range urlValues {
		for _, value := range values {
			p.query.Add(key, value)
		}
	}
	return p
}

func (p *pathBuilder) Build() string {
	pathSegments := strings.Join(append([]string{p.baseURL}, p.pathSegments...), "/")
	fullPath := p.protocol + pathSegments
	if queryParameters := p.query.Encode(); len(queryParameters) > 0 {
		return fullPath + "?" + p.query.Encode()
	}
	p.pathSegments = make([]string, 0)
	p.query = make(url.Values)

	return fullPath
}
