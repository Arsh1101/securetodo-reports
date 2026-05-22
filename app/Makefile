.PHONY: test run tidy build

test:
	go test ./...

run:
	go run ./cmd/web

tidy:
	go mod tidy

build:
	go build -o bin/securetodo ./cmd/web
