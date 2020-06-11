# Heavily based off https://sohlich.github.io/post/go_makefile/
# pared down from the above and added deps download via go mod

# Go parameters
GOCMD=~/.go/bin/go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=referenceProject.out
GETDEPS=$(GOCMD) mod download

all: deps test build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v
		./$(BINARY_NAME)
deps:
		$(GOGETGETDEPS)