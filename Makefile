APP_NAME := go-web-template
PKG := github.com/xbmlz/$(APP_NAME)

VERSION ?= $(shell git describe --match 'v[0-9]*' --dirty='.m' --always --tags)

GO_LDFLAGS ?= -w -X ${PKG}/main.Version=${VERSION}
GO_BUILDTAGS ?=
BUILD_FLAGS ?=
DESTDIR ?= ./bin

ifeq ($(OS),Windows_NT)
    DETECTED_OS = Windows
else
    DETECTED_OS = $(shell uname -s)
endif

ifeq ($(DETECTED_OS),Windows)
	BINARY_EXT=.exe
endif

all: build

.PHONY: build
build: 
	GO111MODULE=on go build $(BUILD_FLAGS) -trimpath -tags "$(GO_BUILDTAGS)" -ldflags "$(GO_LDFLAGS)" -o "$(DESTDIR)/$(APP_NAME)$(BINARY_EXT)" ./cmd
