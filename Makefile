.PHONY: build run test clean

BIN := bin/fs

build:
	@go build -o $(BIN) .

run: build
	@$(BIN)

test:
	@go test ./... -v

clean:
	@rm -rf bin