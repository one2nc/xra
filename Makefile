# Variables.

xra_linux:
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 GOCACHE=/tmp/gocache go build \
		-o xra -a -installsuffix cgo .

test:
	echo "test"
