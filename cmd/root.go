package main

import (
	"os"

	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"xujiajiadexiaokeai.github.com/ebpf-kube-agent/pkg/cmd"
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
