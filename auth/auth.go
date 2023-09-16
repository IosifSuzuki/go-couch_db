package auth

import (
	"encoding/base64"
	"net/http"
)

type Auth interface {
	AddAuth(*http.Request)
}

type auth struct {
	Username string
	Password string
}

func NewAuth(username, password string) Auth {
	return &auth{
		Username: username,
		Password: password,
	}
}

func (a *auth) basicAuth() string {
	auth := []byte(a.Username + ":" + a.Password)
	prefix := "Basic "
	return prefix + base64.StdEncoding.EncodeToString(auth)
}

func (a *auth) AddAuth(req *http.Request) {
	req.Header.Add("Authorization", a.basicAuth())
}
