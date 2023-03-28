package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootLong = `execute, and manage ebpf programs.`
)

func NewRootCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:                   "ekctl",
		DisableFlagsInUseLine: true,
		Short:                 `Execute and manage ebpf programs`,
		Long:                  rootLong,
		Run: func(c *cobra.Command, args []string) {
			cobra.NoArgs(c, args)
			c.Help()
		},
	}

	cmd.AddCommand(NewRunCommand())
	cmd.AddCommand(NewListCommand())

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
