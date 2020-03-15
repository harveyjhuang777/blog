#!make

# Go parameters
BUILD_ENV=CGO_ENABLED=0
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOVET=$(GOCMD) vet
TARGET_EXEC=blog

all: clean set_up test build run

clean: 
	$(GOCLEAN)
	rm -rf build/linux
	rm -rf build/osx
	rm -rf build/windows
	rm -rf build/local

set_up:
	mkdir -p build/linux
	mkdir -p build/osx
	mkdir -p build/windows
	mkdir -p build/local

download:
	$(GOMOD) download

test: 
	$(GOTEST) -v ./...

vrun:
	CGO_ENABLED=0 GOFLAGS=-mod=vendor $(GOBUILD) -o build/local/$(TARGET_EXEC) -v ./cmd/main.go
	./build/local/$(TARGET_EXEC)

vet:
	$(GOVET) -all ./...

build: clean set_up download
	$(GOBUILD) -o build/local/$(TARGET_EXEC) -v ./cmd/main.go

race: clean set_up
	$(GOBUILD) -o build/local/$(TARGET_EXEC) -v -race ./cmd/main.go
	./build/local/$(TARGET_EXEC)

run: build
	./build/local/$(TARGET_EXEC)

traceGC: build
	GODEBUG=gctrace=1 ./build/local/$(TARGET_EXEC)

# Cross compilation
build-linux: set_up
	GOOS=linux $(GOBUILD) -o build/linux/${TARGET_EXEC} -v ./main.go
build-osx: set_up
	GOOS=darwin $(GOBUILD) -o build/osx/${TARGET_EXEC} -v ./main.go
build-windows: setup
	GOOS=windows $(GOBUILD) -o build/windows/${TARGET_EXEC}.exe -v ./main.go
