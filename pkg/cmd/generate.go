package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

var (
	generateShort = "generate"
	generateLong  = `generate bpf programs`
)

// GenerateOptions
type GenerateOptions struct {
	configFlags *genericclioptions.ConfigFlags
	genericclioptions.IOStreams
	namespace    string
	pod          string
	program      string
	clientConfig *rest.Config
}

func NewGenerateOptions(streams genericclioptions.IOStreams) *GenerateOptions {
	return &GenerateOptions{
		configFlags: genericclioptions.NewConfigFlags(false),
		IOStreams:   streams,
	}
}

func NewGenerateCommand(factory cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewGenerateOptions(streams)

	cmd := exec.Command("prog")

	// cmd := &cobra.Command{
	// 	Use:   "gen",
	// 	Short: generateShort,
	// 	Long:  generateLong,
	// 	RunE: func(c *cobra.Command, args []string) error {
	// 		if err := o.Complete(factory); err != nil {
	// 			return err
	// 		}
	// 		if err := o.Run(); err != nil {
	// 			fmt.Fprintln(o.ErrOut, err.Error())
	// 			return nil
	// 		}
	// 		return nil
	// 	},
	// }
	// cmd.Flags().StringVar(&o.namespace, "namespace", o.namespace, "the namespace which the target pod exists")
	// cmd.Flags().StringVar(&o.program, "program", o.program, "program name")
	// return cmd

}
func (o *GenerateOptions) Complete(factory cmdutil.Factory) error {
	var err error
	o.clientConfig, err = factory.ToRESTConfig()
	if err != nil {
		return err
	}

	return nil
}

func (o *GenerateOptions) Run() error {
	// juid := uuid.NewUUID()
	// podName := o.pod
	// namespace := o.namespace
	// program := o.program
	// cmdPrefix := "cd ebpf/examples && go run -exec sudo ./"
	// cmd := cmdPrefix + program
	// clientset, err := kubernetes.NewForConfig(o.clientConfig)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req := clientset.CoreV1().RESTClient().Post().
	// 	Resource("pods").
	// 	Name(podName).
	// 	Namespace(namespace).
	// 	SubResource("exec").
	// 	VersionedParams(&v1.PodExecOptions{
	// 		Command: []string{"sh", "-c", cmd},
	// 		Stdin:   true,
	// 		Stdout:  true,
	// 		Stderr:  true,
	// 		TTY:     true,
	// 	}, scheme.ParameterCodec)
	// executor, err := remotecommand.NewSPDYExecutor(o.clientConfig, "POST", req.URL())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// var stderr bytes.Buffer
	// if err = executor.Stream(remotecommand.StreamOptions{
	// 	Stdin:  os.Stdin,
	// 	Stdout: os.Stdout,
	// 	Stderr: &stderr,
	// 	Tty:    true,
	// }); err != nil {
	// 	fmt.Println(err)
	// }
	return nil
}
