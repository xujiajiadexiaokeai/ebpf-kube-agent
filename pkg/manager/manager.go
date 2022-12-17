package manager

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Config contains the basic manager configuration
type Config struct {
	Host     string
	GrpcPort int
	LogLevel log.Level
}

type tlsConfig struct {
	key    string
	Cert   string
	CaCert string
}

func (c *Config) GprcAddr() string {
	return net.JoinHostPort(c.Host, fmt.Sprintf("%d", c.GrpcPort))
}

func newGrpcServer(tlsConf tlsConfig, log *log.Logger) (*grpc.Server, error) {
	s := grpc.NewServer()
	return s, nil
}

type Manager struct {
	grpcServer *grpc.Server
}
