# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Name of your Go application
APPNAME = app

# Binary output directory
BINDIR = ./

# Source code directory
SRCDIR = .

# Main Go source file
MAINFILE = .

# Build target
all: clean build

build:
	$(GOBUILD) -o $(APPNAME) $(MAINFILE)


clean:
	$(GOCLEAN)
	rm -rf $(BINDIR)

test:
	$(GOTEST) ./...

run:
	./app

build-run: build
	run

.PHONY: all build clean test run
