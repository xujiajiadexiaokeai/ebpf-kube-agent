package common

import (
	"context"
	"net"
	"time"

	"github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/manager/pb"
	"google.golang.org/grpc"
)

const DefaultRPCTimeout = 60 * time.Second

var RPCTimeout = DefaultRPCTimeout

func NewGRPClient() (*GrpcClient, error) {
	var ip string
	var port string
	builder := builder(ip, port).WithDefaultTimeout()
	cc, err := builder.build()
	if err != nil {
		return nil, err
	}
	return &GrpcClient{
		ManagerClient: pb.NewManagerClient(cc),
		conn:          cc,
	}, nil

}

type GrpcClient struct {
	pb.ManagerClient
	conn *grpc.ClientConn
}

func (c *GrpcClient) Close() error {
	return c.conn.Close()
}

type grpcBuilder struct {
	options []grpc.DialOption
	address string
	port    string
}

func builder(address string, port string) *grpcBuilder {
	return &grpcBuilder{options: []grpc.DialOption{}, address: address, port: port}
}

func (it *grpcBuilder) WithDefaultTimeout() *grpcBuilder {
	it.options = append(it.options, grpc.WithUnaryInterceptor(TimeoutClientInterceptor(DefaultRPCTimeout)))
	return it
}

// TimeoutClientInterceptor wraps the RPC with a timeout.
func TimeoutClientInterceptor(timeout time.Duration) func(context.Context, string, interface{}, interface{},
	*grpc.ClientConn, grpc.UnaryInvoker, ...grpc.CallOption) error {
	return func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (it *grpcBuilder) build() (*grpc.ClientConn, error) {
	return grpc.Dial(net.JoinHostPort(it.address, it.port), it.options...)
}
