package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

var (
	rootLong = `configure, execute, and manage bpftrace programs.`
)

// RootOptions
type RootOptions struct {
	configFlags *genericclioptions.ConfigFlags

	genericclioptions.IOStreams
}

func NewRootOptions(streams genericclioptions.IOStreams) *RootOptions {
	return &RootOptions{
		configFlags: genericclioptions.NewConfigFlags(false),

		IOStreams: streams,
	}
}

func NewRootCommand(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewRootOptions(streams)

	cmd := &cobra.Command{
		Use:                   "kubectl-ebpf",
		DisableFlagsInUseLine: true,
		Short:                 `Execute and manage bpftrace programs`,
		Long:                  rootLong,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.SetOutput(streams.ErrOut)
		},
		Run: func(c *cobra.Command, args []string) {
			cobra.NoArgs(c, args)
			c.Help()
		},
	}

	matchVersionFlags := cmdutil.NewMatchVersionFlags(o.configFlags)

	f := cmdutil.NewFactory(matchVersionFlags)

	cmd.AddCommand(NewRunCommand(f, streams))
	cmd.AddCommand(NewGenerateCommand(f, streams))

	walk(cmd, func(c *cobra.Command) {
		c.Flags().BoolP("help", "h", false, fmt.Sprintf("Help for the %s command", c.Name()))
	})

	return cmd
}

func walk(c *cobra.Command, f func(*cobra.Command)) {
	f(c)
	for _, c := range c.Commands() {
		walk(c, f)
	}
}
