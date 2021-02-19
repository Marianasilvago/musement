all: build test clean

build: build-amusement

build-amusement-api:
	go build -v -i -o buildtest/amusement .

test:
	go test -v ./pkg/envutils
	go test -v ./internal/version

clean:
	rm -rf buildtest

.PHONY: build build-amusement test clean