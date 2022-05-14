ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif

.DEFAULT_GOAL := docker-push-image

# local dev
test:
	go test ./...
.PHONY:test

build: test
	go build main.go
.PHONY:build

# clear local dev
clear:
	rm main.exe
.PHONY:clear

# docker
docker-build-image:
	docker build -t stakkato95/lambda-service-analytics:latest . -f Dockerfile
.PHONY:docker-build-image

docker-push-image: docker-build-image
	docker push stakkato95/lambda-service-analytics:latest
.PHONY:docker-push-image
