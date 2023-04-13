package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	cm "github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/ekctl/common"
	"github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/manager/pb"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all items",
		Long:  "List all items in the system",
	}
	cmd.AddCommand(NewListPodCommand())
	cmd.AddCommand(NewListProgCommand())
	return cmd
}

func NewListPodCommand() *cobra.Command {
	var namespace string
	cmd := &cobra.Command{
		Use:   "pod",
		Short: "List all pods",
		Long:  "List all pods in the Kubernetes cluster",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing all pods...")
			runListPod(namespace)
		},
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "Specify the namespace to list pods from")

	return cmd
}

func runListPod(namespace string) {
	client, err := cm.NewGRPClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()
	resp, err := client.ListPod(context.Background(), &pb.ListPodRequest{Namespace: namespace})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, pod := range resp.PodList {
		fmt.Printf("Pod: %s, Namespace: %s", pod.Name, pod.Namespace)
	}
}

func NewListProgCommand() *cobra.Command {
	var namespace, pod string
	cmd := &cobra.Command{
		Use:   "prog",
		Short: "List all eBPF programs on a pod",
		Long:  "List all eBPF programs on a pod in the Kubernetes cluster",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Listing all eBPF programs on pod %s in namespace %s...\n", pod, namespace)
			// Add code to list all eBPF programs on pod here
		},
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "Specify the namespace to list pods from")
	cmd.Flags().StringVar(&pod, "pod", "", "Specify the pod to list eBPF programs from")

	return cmd
}
