ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif

.DEFAULT_GOAL := run-with-env

# local dev
test:
	go test ./...
.PHONY:test

build: test
	go build main.go
.PHONY:build

run-with-env: build
	$$env:SERVER_PORT='8080'; $$env:KAFKA_SERVICE='kafka.default.svc.cluster.local:9092'; ./main
.PHONY:run-with-env

# clear local dev
clear:
	rm main.exe
.PHONY:clear

# docker
docker-build-image:
	docker build -t stakkato95/lambda-ingress:latest . -f Dockerfile
.PHONY:docker-build-image

docker-push-image: docker-build-image
	docker push stakkato95/lambda-ingress:latest
.PHONY:docker-push-image

docker-run-tmp-container:
	docker run --rm -p 8080:8080 -d stakkato95/lambda-ingress
.PHONY:docker-local-container
