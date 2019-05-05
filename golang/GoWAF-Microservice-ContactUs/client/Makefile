
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker

all: plugins build test

plugins:
	go build -buildmode=plugin -o plugins/middlewares/method.so plugins/middlewares/method.go

build:
	go build -o webclient main.go
    
test:
	go test -v ./... -cover

docker:
	docker build . --rm --squash -t webclient:latest

run:
	docker run -p 8080:3001 -it --rm --name webclient-run webclient:latest