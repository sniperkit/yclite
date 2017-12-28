# Reference: <https://sohlich.github.io/post/go_makefile/>
GOCMD   := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST  := $(GOCMD) test
GOGET   := $(GOCMD) get
GORUN 	:= $(GOCMD) run

BINARY_NAME := yclite
BINARY_UNIX := $(BINARY_NAME)_unix

all: test build
build: ; $(GOBUILD) -o $(BINARY_NAME) -v
test:  ; $(GOTEST) -v ./...
clean: ; $(GOCLEAN) && rm -f $(BINARY_NAME) $(BINARY_UNIX)
run:   ; $(GORUN) $(BINARY_NAME) -v ./...

deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop


# Cross compilation
build-linux:  ; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build: ; docker build -t yclite:0.2 .