# Variables.

DEP := $(shell command -v dep 2> /dev/null)

deps:
ifndef DEP
	dep version || (curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh)
	dep ensure -v
endif
	dep ensure -v

xra_linux: deps
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 GOCACHE=/tmp/gocache go build \
		-o xra -a -installsuffix cgo .

test: deps
	go test -v ./...
