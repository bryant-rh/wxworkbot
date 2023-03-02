PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))
VERSION = $(shell git describe --abbrev=0 --tags|awk -F 'v' '{print $$2}')
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

GOBUILD=CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -ldflags "-X $(PKG)/cmd/$(NAME).version=$(Version)"
PLATFORM := linux/amd64,linux/arm64

WORKSPACE ?= name

clean:
	rm -rf ./cmd/$(WORKSPACE)/out

upgrade:
	go get -u ./...

tidy:
	go mod tidy

build: 
	$(GOBUILD)

docker:
	docker buildx build --push --progress plain --platform=${PLATFORM}	\
		--cache-from "type=local,src=/tmp/.buildx-cache" \
		--cache-to "type=local,dest=/tmp/.buildx-cache" \
		--file=./Dockerfile \
		--tag=bryantrh/wxworkbot:${VERSION} \
		.