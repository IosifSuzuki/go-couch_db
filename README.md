# go-couch_db

Package that helps interact with [CouchDB](https://couchdb.apache.org). The repository is in active development mode. Please feel free to make a pull request.

## Example connect to [CouchDB](https://couchdb.apache.org)

```go
conn, err := connection.NewConnection(client.HttpProtocol, host, port, username, password, logger.DebugLevel)
if err != nil {
	log.Fatalf("can't create connection: %v", err)
}
ok, err := conn.Ping()
if err != nil {
	log.Errorf("pinged to db: %v", err)
} else {
	log.Debugf("pinged to db: %t", ok)
}
```

## Install

go get -u github.com/IosifSuzuki/go-couch_db