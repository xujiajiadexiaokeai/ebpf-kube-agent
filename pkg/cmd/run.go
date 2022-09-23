package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"xujiajiadexiaokeai.github.com/ebpf-kube-agent/pkg/agent"
	"xujiajiadexiaokeai.github.com/ebpf-kube-agent/pkg/ebpfjob"
)

var (
	runShort = "run"
	runLong  = `execute bpf programs`
)

// RunOptions
type RunOptions struct {
	configFlags *genericclioptions.ConfigFlags
	genericclioptions.IOStreams
	namespace       string
	targetNamespace string
	pod             string
	container       string
	clientConfig    *rest.Config
}

func NewRunOptions(streams genericclioptions.IOStreams) *RunOptions {
	return &RunOptions{
		configFlags: genericclioptions.NewConfigFlags(false),
		IOStreams:   streams,
	}
}

func NewRunCommand(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewRunOptions(streams)

	cmd := &cobra.Command{
		Use:   "run",
		Short: runShort,
		Long:  runLong,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.SetOutput(streams.ErrOut)
		},
		Run: func(c *cobra.Command, args []string) {
			cobra.NoArgs(c, args)
			c.Help()
		},
	}

	flags := cmd.PersistentFlags()
	o.configFlags.AddFlags(flags)

	// matchVersionFlags := cmdutil.NewMatchVersionFlags(o.configFlags)
	// matchVersionFlags.AddFlags(flags)

	// f := cmdutil.NewFactory(matchVersionFlags)

	// cmd.AddCommand()

	return cmd

}

func (o *RunOptions) Run() error {
	juid := uuid.NewUUID()

	clientset, err := kubernetes.NewForConfig(o.clientConfig)
	if err != nil {
		return err
	}

	target, err := ebpfjob.GetJobTarget(clientset, o.pod, o.targetNamespace)
	if err != nil {
		return err
	}
	jc := ebpfjob.NewJobClient(clientset, o.namespace)

	ej := ebpfjob.EbpfJob{
		ID: juid,
	}

	job, err := jc.CreateJob(ej)
	if err != nil {
		return err
	}

	fmt.Fprintf(o.IOStreams.Out, "job %s created\n", ej.ID)

	a := agent.NewAgent()

	a.AttachJob(job)
}
