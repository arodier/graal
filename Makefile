


# Graal:
# go remote administration api for linux

# Standard env
GO ?= /usr/bin/go
LIB_OPTIONS ?= -linkshared
BUILD_OPTIONS ?= -race

all: build

clean:
	rm -rf lib bin

# Install go standard librairies as shared
shared-install:
	@echo -n 'Installing standard librairies (as shared): '
	@sudo $(GO) install -buildmode=shared std && echo 'OK' || echo 'Fail'

build:
	$(GO) build $(BUILD_OPTIONS) -o bin/graal src/main.go

run:
	./bin/graal
