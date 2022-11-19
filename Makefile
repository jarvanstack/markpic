# config
bindir = ./bin
releaseDir = ./release
mainFile = ./main.go

# var
make_dir:=$(shell pwd)
app_name:=$(shell basename $(make_dir))
version=$(shell git describe --tags --always)

## build: Builds the project
.PHONY: build
build: 
	GOOS=linux GOARCH=amd64 go build -o $(bindir)/$(app_name).linux $(mainFile)
	GOOS=darwin GOARCH=amd64 go build -o $(bindir)/$(app_name).darwin $(mainFile)
	GOOS=windows GOARCH=amd64 go build -o $(bindir)/$(app_name).exe $(mainFile)

## release: Builds the project for release	
.PHONY: release
release: clean build
	tar -zcvf $(releaseDir)/$(app_name)-linux-$(version).tar.gz $(bindir)/$(app_name).linux 
	tar -zcvf $(releaseDir)/$(app_name)-darwin-$(version).tar.gz $(bindir)/$(app_name).darwin 
	tar -zcvf $(releaseDir)/$(app_name)-exe-$(version).tar.gz $(bindir)/$(app_name).exe 

## clean: Cleans the project
.PHONY: clean
clean:
	rm -rf $(bindir)/*
	rm -rf $(releaseDir)/*
	go clean

## help: Show this help info.
.PHONY: help
help: Makefile
	@printf "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS