package client

import (
	"github.com/IosifSuzuki/go-couch_db/auth"
	"github.com/IosifSuzuki/go-couch_db/logger"
	"io"
	"moul.io/http2curl"
	"net/http"
	"time"
)

const (
	contentType string = "Content-Type"
)

type HttpClient struct {
	client *http.Client
	auth   auth.Auth
	log    logger.Logger
}

func NewHttpClient(auth auth.Auth, log logger.Logger) HttpClient {
	return HttpClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		auth: auth,
		log:  log,
	}
}

func (h *HttpClient) Execute(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		h.log.Errorf("create new request with error %v", err)
		return nil, err
	}
	curl, err := http2curl.GetCurlCommand(req)
	if err != nil {
		h.log.Errorf("create representable curl command from request has finished with error: %v", err)
	}
	h.log.Debug(curl.String())
	h.auth.AddAuth(req)
	req.Header.Set(contentType, "application/json")
	return h.client.Do(req)
}
