.PHONY: build test clean

build:
	go build -o bin/database main.go

run:
	go run main.go

test:
	go test ./...

clean:
	rm -rf bin/