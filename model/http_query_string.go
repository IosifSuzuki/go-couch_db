package model

import "net/url"

type HttpQueryString interface {
	BuildUrlValues() (url.Values, error)
}
