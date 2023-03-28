package manager

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/manager/pb"
)

func (s *DaemonServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logrus.Info("Received: ", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName() + ", This is manager."}, nil
}

func (s *DaemonServer) GetKubernetesVersion(ctx context.Context, in *pb.KubernetesVersionRequest) (*pb.KubernetesVersionReply, error) {
	kubernetesVersion, err := s.provider.GetKubernetesVersion()
	if err != nil {
		logrus.Error(err)
	}
	return &pb.KubernetesVersionReply{Version: kubernetesVersion}, nil
}

func (s *DaemonServer) ListPod(ctx context.Context, in *pb.ListPodRequest) (*pb.ListPodResponse, error) {
	pods, err := s.provider.ListPod(ctx, in.GetNamespace())
	if err != nil {
		logrus.Error(err)
	}
	podList := []*pb.Pod{}
	for _, pod := range pods {
		podList = append(podList, &pb.Pod{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		})
	}

	return &pb.ListPodResponse{PodList: podList}, nil
}
