# veligroxy
A gopher-friendly velib proxy
=======

## Creating a module

```
$ go mod init github.com/fagossa/golang-rest-api
```


===================
TODO
```
go mod tidy
go mod download
```

GO111MODULE=on go mod init
GO111MODULE=on go mod tidy
GO111MODULE=on go mod download
GO111MODULE=on go mod vendor

GOPROXY=https://gocenter.io go mod tidy

===================

## Building (fetching dependencies)

### Rebuild all

```
$ go build ./...
```

### Rebuild using main
```
go build -o bin/server cmd/api/main.go
```

## Running

```
$ PORT=8080 DIAG_PORT=8081 ./bin/server
```

## Shutdown
kill -SIGTERM ${process id}


# service template

https://github.com/takama/caldera
