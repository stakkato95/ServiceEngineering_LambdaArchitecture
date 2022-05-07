ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif

.DEFAULT_GOAL := run-with-env

test:
	go test ./...
.PHONY:test

build: test
	go build main.go
.PHONY:build

run-with-env: build
	$$env:SERVER_ADDRESS='localhost';$$env:SERVER_PORT='8080'; ./main
.PHONY:run-with-env

clear:
	rm main.exe
.PHONY:clear