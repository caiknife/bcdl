.PHONY: build test

build:
	go mod tidy
	go build -ldflags="-s -w" -tags=jsoniter -o ./out/bcdl ./bcdl

test:
	go test ./...
