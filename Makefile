all: fmt run

fmt:
	go fmt ./...

staticcheck:
	staticcheck ./...

run-lab1:
	go run lab1/cmd/main.go

run-lab3:
	go run lab3/cmd/main.go

.PHONY: all fmt staticcheck run
