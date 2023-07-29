# Makefile for building the crosscheck CLI tool

# Targets
.PHONY: all build clean

all: build

build:
	@go build -o ~/go/bin/crosssight

clean:
	@rm -rf ~/go/bin/crosssight
