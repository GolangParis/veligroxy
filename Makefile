PROJECT?=github.com/GolangParis/veligroxy

RELEASE?=0.0.0
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	CGO_ENABLED=0 go build \
		-ldflags "-s -w -X ${PROJECT}/internal/version.Version=${RELEASE} \
		-X ${PROJECT}/internal/version.Commit=${COMMIT} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
		-o bin/server ${PROJECT}/cmd/api

test:
	GO111MODULE=on go test ./... -v

run:
		PORT=8080 DIAG_PORT=8081 ./bin/server
