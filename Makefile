# Variables.

DEP := $(shell command -v dep 2> /dev/null)

deps:
ifndef DEP
	go get -u github.com/golang/dep/cmd/dep
endif
	dep ensure -v

xra_linux: deps
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 GOCACHE=/tmp/gocache go build \
		-o xra -a -installsuffix cgo .

test: deps
	go test -v ./...
