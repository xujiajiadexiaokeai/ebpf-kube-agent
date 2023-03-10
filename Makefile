
agent ?= bin/ebpf-agent
manager ?= bin/ebpf-manager
kubectl-ebpf ?= bin/kubectl-ebpf 
GO ?= go

.PHONY: build
build: clean
	$(GO) build -o ${agent} ./cmd/agent
	$(GO) build -o ${manager} ./cmd/manager
	$(GO) build -o ${kubectl-ebpf} ./cmd/root.go

.PHONY: image-build
image-build:
	docker buildx build \
	-f images/manager/Dockerfile \
	-t ebpf-manager:0.1 \
	.
.PHONY: kind-image-load
kind-image-load:
	kind load docker-image ebpf-manager:0.1

.PHONY: deploy-manager
deploy-manager:
	kubectl apply -f k8s/manager/manager-deployment.yaml
	kubectl apply -f k8s/manager/manager-service.yaml

.PHONY: clean
clean: 
	rm -rf bin
