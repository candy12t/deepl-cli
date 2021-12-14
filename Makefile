OBJ = bin/deepl-cli
VERSION := $(shell git describe --tags --abbrev=0)

all: clean build

.PHONY: build
build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(OBJ) -v ./cmd/deepl-cli

.PHONY: clean
clean:
	rm -rf $(OBJ)

.PHONY: test
test:
	go test ./... -count=1
