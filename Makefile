
agent ?= bin/agent
kubectl-ebpf ?= bin/kubectl-ebpf 
GO ?= go

.PHONY: build
build: clean
	$(GO) build -o ${agent} ./cmd/agent
	$(GO) build -o ${kubectl-ebpf} ./cmd/root.go

.PHONY: clean
clean: 
	rm -rf bin
