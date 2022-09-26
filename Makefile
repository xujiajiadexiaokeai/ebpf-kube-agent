
agent ?= bin/agent
GO ?= go




.PHONY: build
build: clean ${agent}

${agent}:
	$(GO) build -o $@ ./cmd/agent