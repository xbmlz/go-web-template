PKG := github.com/xbmlz/go-web-template

VERSION ?= $(shell git describe --match 'v[0-9]*' --dirty='.m' --always --tags)

GO_LDFLAGS ?= -w -X ${PKG}/internal.Version=${VERSION}
GO_BUILDTAGS ?=
BUILD_FLAGS ?=
DESTDIR ?=

ifeq ($(DETECTED_OS),Windows)
	BINARY_EXT=.exe
endif

all: build

.PHONY: build
build: 
	GO111MODULE=on go build $(BUILD_FLAGS) -trimpath -tags "$(GO_BUILDTAGS)" -ldflags "$(GO_LDFLAGS)" -o "$(or $(DESTDIR),./bin/build)/docker-compose$(BINARY_EXT)" ./cmd
