#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export GOOS=linux
export GOARCH=amd64
export DEBUG=true
export SERVER_ADDRESS=0.0.0.0:8080
export APP=elastest

export LDFLAGS="-w -s"

all: format lint build

build:
	go build .

build-static:
	CGO_ENABLED=0 go build -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	go run .

run-mac:
	GOOS=darwin go run .

############################################################
# Format & Lint
############################################################

format:
	which goimports || GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R goimports -w R
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R gofmt -s -w R

check-gometalinter:
	which golangci-lint || (GO111MODULE=off go get -u -v github.com/golangci/golangci-lint)


fast-lint: check-gometalinter
	golangci-lint run --enable-all --deadline 10m $(ROOT)/...


############################################################
# Test
############################################################

test:
	go test -v ./...

ci-test:
	go test -race $(go list ./... | grep -v /vendor/)
	go test ./... -v -coverprofile .coverage.out
	go tool cover -func .coverage.out

.PHONY: build get install run watch start stop restart clean up ci-test
