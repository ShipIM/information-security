all: fmt run

fmt:
	go fmt ./...

staticcheck:
	staticcheck ./...

run:
	go run cmd/main.go

.PHONY: all fmt staticcheck run
