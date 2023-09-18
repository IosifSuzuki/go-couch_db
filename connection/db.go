package connection

import (
	"github.com/IosifSuzuki/go-couch_db/auth"
	"github.com/IosifSuzuki/go-couch_db/client"
	"github.com/IosifSuzuki/go-couch_db/logger"
	"net"
)

type connection struct {
	db          string
	pathBuilder client.PathBuilder
	httpClient  client.HttpClient
	log         logger.Logger
}

func NewConnection(protocol, db, host, port, username, password string, logLevel logger.LogLevel) (Operation, error) {
	baseAuth := auth.NewAuth(username, password)
	log := logger.NewLogger(logLevel)
	httpClient := client.NewHttpClient(baseAuth, log)
	pathBuilder := client.NewPath(protocol, net.JoinHostPort(host, port))
	return &connection{
		db:          db,
		pathBuilder: pathBuilder,
		httpClient:  httpClient,
		log:         log,
	}, nil
}
