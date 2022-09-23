package ebpfjob

import (
	"io"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	batchv1typed "k8s.io/client-go/kubernetes/typed/batch/v1"
	corev1typed "k8s.io/client-go/kubernetes/typed/core/v1"

	batchv1 "k8s.io/api/batch/v1"
)

type EbpfJobClient struct {
	JobClient    batchv1typed.JobInterface
	ConfigClient corev1typed.ConfigMapInterface
	outStream    io.Writer
}

type EbpfJob struct {
	Name        string
	ID          types.UID
	Namespace   string
	Output      string
	Program     string
	ProgramArgs string
	StartTime   *metav1.Time
	Status      EbpfJobStatus
}

type EbpfJobStatus string

func NewJobClient(clientset kubernetes.Interface, namespace string) *EbpfJobClient {
	return &EbpfJobClient{
		JobClient:    clientset.BatchV1().Jobs(namespace),
		ConfigClient: clientset.CoreV1().ConfigMaps(namespace),
	}
}

func (c *EbpfJobClient) CreateJob(ej EbpfJob) (*batchv1.Job, error) {
	job := ej.Job()
	return job, nil
}

func (ej *EbpfJob) Job() *batchv1.Job {
	job := &batchv1.Job{}
	return job
}
