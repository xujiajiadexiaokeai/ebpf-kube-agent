[WIP]ebpf-kube-agent
---

`ebpf-kube-agent` is a tool that uses [cilium/ebpf](https://github.com/cilium/ebpf) as an ebpf program deployer to deploy ebpf programs on Kubernetes nodes and collect data.

# Features
- [X] Support deploying ebpf programs on Kubernetes nodes
- [ ] Support collecting data from ebpf maps
- [ ] Support sending data to Prometheus
- [ ] Support customizing ebpf programs and maps

# Design
The tool consists of three parts: a CLI client, a manager and an agents group.

The CLI client is used to interact with the manager, such as deploying ebpf programs, querying data, etc.

The manager is deployed as a Deployment with one replica. It acts as a gRPC server that receives data from agents and exposes them as Prometheus metrics.

The agent is deployed as a DaemonSet on each node. It is responsible for loading and attaching ebpf programs to hooks, creating and managing ebpf maps, collecting data from maps periodically, and sending data to manager via gRPC.

# Usage

## Prerequisites
A Kubernetes cluster with version >= 1.18

Linux kernel version >= 5.4 with BTF enabled

Go version >= 1.16

clang/llvm version >= 11

## Installation
**TODO**