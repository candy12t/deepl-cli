OBJ = bin/deepl-cli
VERSION := $(shell git describe --tags --abbrev=0)

all: test clean build

.PHONY: build
build:
	CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X github.com/candy12t/deepl-cli/internal/build.Version=$(VERSION)" -o $(OBJ) -v ./cmd/deepl-cli

.PHONY: clean
clean:
	rm -rf $(OBJ)

.PHONY: test
test:
	go test -race ./... -count=1
