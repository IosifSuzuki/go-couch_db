package connection

import "net/http"

type Operation interface {
	Ping() (bool, error)
}

func (c *connection) Ping() (bool, error) {
	path := c.pathBuilder.Build()
	resp, err := c.httpClient.Execute(http.MethodHead, path, nil)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == http.StatusOK, nil
}
