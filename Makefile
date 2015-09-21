
# Standard env
GO ?= /usr/bin/go
MARKDOWN ?= /usr/bin/markdown
LIB_OPTIONS ?= -linkshared
BUILD_OPTIONS ?= -race

all: clean build

clean:
	rm -rf lib bin

# Install go standard librairies as shared
shared-install:
	@echo -n 'Installing standard librairies (as shared): '
	@sudo $(GO) install -buildmode=shared std && echo 'OK' || echo 'Fail'

build:
	$(GO) build $(BUILD_OPTIONS) -o bin/graal src/main.go

# For now, the 'home page' is generated using the README.md,
# until we'll use the source code to generate docs
docs:
	$(MARKDOWN) README.md >docs/home.html

run:
	./bin/graal --home=./docs/home.html
