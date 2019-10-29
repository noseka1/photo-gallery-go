PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

default: build

clean:
	rm -f ./bin/${PROJECTNAME}

mod:
	go mod download

build:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit

start:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit
	./bin/$(PROJECTNAME)

build_static:
	CGO_ENABLED=0 go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go

image: build_static
	podman build -t $(PROJECTNAME) .
