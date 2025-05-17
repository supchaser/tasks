APP      := echo-server
VERSION  ?= latest
REGISTRY ?= docker.io
REPO     ?= supchaser/cloudru-camp-echo-server

.PHONY: all build docker-build docker-push run

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -v -ldflags="-s -w" -o bin/$(APP) ./cmd/main

docker-build: build
	docker build -t $(REGISTRY)/$(REPO):$(VERSION) .

docker-push: docker-build
	docker push $(REGISTRY)/$(REPO):$(VERSION)

run:
	docker run --rm -p 8000:8000 \
		-e AUTHOR="Valentin Stremin (supchaser)" \
		$(REGISTRY)/$(REPO):$(VERSION)
