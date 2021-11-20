OBJ = bin/deepl-cli
all: clean build

build:
	go build -o $(OBJ) ./cmd/deepl-cli/main.go

clean:
	rm -rf $(OBJ)
