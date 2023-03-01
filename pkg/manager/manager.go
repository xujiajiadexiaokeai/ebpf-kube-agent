package manager

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/manager/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedManagerServer
}

// Config contains the basic manager configuration
type Config struct {
	Host     string
	GrpcPort int
	LogLevel log.Level

	tlsConfig
}

type tlsConfig struct {
	key    string
	Cert   string
	CaCert string
}

func (c *Config) GprcAddr() string {
	return net.JoinHostPort(c.Host, fmt.Sprintf("%d", c.GrpcPort))
}

type DaemonServer struct {
	provider *Provider
}

func newDaemonServer() (*DaemonServer, error) {
	provider, err := getKubernetesProvider()
	if err != nil {
		return nil, err
	}
	return &DaemonServer{
		provider: provider,
	}, nil
}

func newGrpcServer(tlsConf tlsConfig) (*grpc.Server, error) {
	s := grpc.NewServer()
	pb.RegisterManagerServer(s, &Server{})
	return s, nil
}

type Manager struct {
	grpcServer   *grpc.Server
	daemonServer *DaemonServer
	conf         *Config
	logger       log.Logger
}

func BuildManager(conf *Config, log log.Logger) (*Manager, error) {
	manager := &Manager{conf: conf, logger: log}
	var err error
	manager.grpcServer, err = newGrpcServer(manager.conf.tlsConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create grpc server")
	}
	manager.daemonServer, err = newDaemonServer()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create daemon server")
	}
	return manager, nil
}

func (m *Manager) Start() error {
	m.logger.Info("starting manager")
	grpcBindAddr := m.conf.GprcAddr()
	m.logger.Info("starting grpc server on", grpcBindAddr)
	grpcLintener, err := net.Listen("tcp", grpcBindAddr)
	if err != nil {
		return errors.Wrap(err, "failed to listen on grpc bind address")
	}
	var eg errgroup.Group

	eg.Go(func() error {
		if err := m.grpcServer.Serve(grpcLintener); err != nil {
			return errors.Wrap(err, "failed to serve grpc server")
		}
		return nil
	})
	return eg.Wait()
}

func (m *Manager) Stop() error {
	m.grpcServer.GracefulStop()
	return nil
}

func getKubernetesProvider() (*Provider, error) {
	return nil, nil
}
