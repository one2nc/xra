# Variables.

DEP := $(shell command -v dep 2> /dev/null)

deps:
ifndef DEP
	go get -u github.com/golang/dep/cmd/dep
endif
	dep ensure -v

xra_linux: deps
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
							-a -ldflags '-w -extldflags "-static"' .

test: deps
	go test -v ./...
