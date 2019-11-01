PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

default: build

clean:
	rm -f ./bin/${PROJECTNAME}

mod:
	go mod download

build:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go

start:
	./bin/$(PROJECTNAME)

build_static: export CGO_ENABLED=0
build_static: build

image:
	buildah build-using-dockerfile -t $(PROJECTNAME) .
