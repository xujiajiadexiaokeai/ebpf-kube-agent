package manager

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/manager/pb"
)

type server struct {
	pb.UnimplementedManagerServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logrus.Info("Received: ", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName() + ", This is manager."}, nil
}
