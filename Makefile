all: fmt run

fmt:
	go fmt ./...

staticcheck:
	staticcheck ./...

run-lab1:
	go run lab1/cmd/main.go

run-lab2:
	go run lab2/cmd/main.go

run-lab3:
	go run lab3/cmd/main.go

run-lab4:
	go run lab4/cmd/main.go

run-lab5:
	go run lab5/cmd/main.go

.PHONY: all fmt staticcheck run
