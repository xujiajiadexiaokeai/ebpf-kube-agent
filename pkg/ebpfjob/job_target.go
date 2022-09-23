package ebpfjob

import (
	"k8s.io/client-go/kubernetes"
)

type JobTarget struct {
	Pod string
}

func GetJobTarget(clientset kubernetes.Interface, pod, targetNamespace string) (*JobTarget, error) {
	target := JobTarget{}
	return &target, nil
}
