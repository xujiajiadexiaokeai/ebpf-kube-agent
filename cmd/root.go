package main

import (
	"os"

	"github.com/spf13/pflag"
	"github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/ekctl/cmd"
)

func main() {
	flags := pflag.NewFlagSet("ekctl", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewRootCommand()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
