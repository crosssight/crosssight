# Makefile for building the crosscheck CLI tool

# Targets
.PHONY: all build clean

all: build

build:
	@go build -o /Users/jonathanpick/go/bin/crosscheck

clean:
	@rm -rf /Users/jonathanpick/go/bin/crosscheck
