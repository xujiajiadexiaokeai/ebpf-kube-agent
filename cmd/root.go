package main

import (
	"os"

	"github.com/spf13/pflag"
	"github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-ebpf", pflag.ExitOnError)
	pflag.CommandLine = flags

	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	root := cmd.NewRootCommand(streams)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
