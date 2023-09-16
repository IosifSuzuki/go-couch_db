package connection

import (
	"github.com/IosifSuzuki/go-couch_db/auth"
	"github.com/IosifSuzuki/go-couch_db/client"
	"github.com/IosifSuzuki/go-couch_db/logger"
	"net"
)

type connection struct {
	pathBuilder client.PathBuilder
	httpClient  client.HttpClient
}

func NewConnection(protocol, host, port, username, password string, logLevel logger.LogLevel) (Operation, error) {
	baseAuth := auth.NewAuth(username, password)
	log := logger.NewLogger(logLevel)
	httpClient := client.NewHttpClient(baseAuth, log)
	pathBuilder := client.NewPath(protocol, net.JoinHostPort(host, port))
	return &connection{
		pathBuilder: pathBuilder,
		httpClient:  httpClient,
	}, nil
}
