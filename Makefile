.PHONY: build clean

build:
	go build -o /usr/bin/tammy ./cmd/main.go
