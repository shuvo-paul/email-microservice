.PHONY: build test run clean

build:
	go build -o ./tmp/main ./cmd/server/main.go

test:
	go test ./..

run:
	./tmp/main

clean:
	rm -rf ./tmp
