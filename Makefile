OBJ = bin/deepl-cli
VERSION := $(shell git describe --tags --abbrev=0)

all: clean build

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(OBJ) -v ./cmd/deepl-cli

clean:
	rm -rf $(OBJ)
