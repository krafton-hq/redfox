VERSION=0.0.0-makefile-build

# Parameter: Version
build:
	CGO_ENABLED=0 \
	go build -ldflags='-s -w -X "github.com/krafton-hq/red-fox/server/application.version=$(VERSION)"' -trimpath
