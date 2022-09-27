package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/scheme"
)

var (
	runShort = "run"
	runLong  = `execute bpf programs`
)

// RunOptions
type RunOptions struct {
	configFlags *genericclioptions.ConfigFlags
	genericclioptions.IOStreams
	namespace    string
	pod          string
	program      string
	clientConfig *rest.Config
}

func NewRunOptions(streams genericclioptions.IOStreams) *RunOptions {
	return &RunOptions{
		configFlags: genericclioptions.NewConfigFlags(false),
		IOStreams:   streams,
	}
}

func NewRunCommand(factory cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewRunOptions(streams)

	cmd := &cobra.Command{
		Use:   "run",
		Short: runShort,
		Long:  runLong,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(factory); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				fmt.Fprintln(o.ErrOut, err.Error())
				return nil
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&o.namespace, "namespace", o.namespace, "the namespace which the target pod exists")
	cmd.Flags().StringVar(&o.pod, "pod", o.pod, "the pod which the program to execute")
	cmd.Flags().StringVar(&o.program, "program", o.program, "program name")
	return cmd

}
func (o *RunOptions) Complete(factory cmdutil.Factory) error {
	var err error
	o.clientConfig, err = factory.ToRESTConfig()
	if err != nil {
		return err
	}

	return nil
}

func (o *RunOptions) Run() error {
	// juid := uuid.NewUUID()
	podName := o.pod
	namespace := o.namespace
	program := o.program
	cmdPrefix := "cd ebpf/examples && go run -exec sudo ./"
	cmd := cmdPrefix + program
	clientset, err := kubernetes.NewForConfig(o.clientConfig)
	if err != nil {
		fmt.Println(err)
	}
	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Command: []string{"sh", "-c", cmd},
			Stdin:   true,
			Stdout:  true,
			Stderr:  true,
			TTY:     true,
		}, scheme.ParameterCodec)
	executor, err := remotecommand.NewSPDYExecutor(o.clientConfig, "POST", req.URL())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("pod: %s, namespace: %s, function name: %s", podName, namespace, program))
	fmt.Println("program start")
	var stderr bytes.Buffer
	if err = executor.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: &stderr,
		Tty:    true,
	}); err != nil {
		fmt.Println(err)
	}
	return nil
}
