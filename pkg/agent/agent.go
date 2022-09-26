package agent

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"xujiajiadexiaokeai.github.com/ebpf-kube-agent/pkg/log"
)

type Config struct {
	Host     string
	HTTPPort int
}

func (c *Config) HttpAddr() string {
	return net.JoinHostPort(c.Host, fmt.Sprintf("%d", c.HTTPPort))
}

type DaemonServer struct {
	rootLogger logr.Logger
}

func (s *DaemonServer) getLoggerFromContext(ctx context.Context) logr.Logger {
	return log.EnrichLoggerWithContext(ctx, s.rootLogger)
}

func newDaemonServer(log logr.Logger) (*DaemonServer, error) {
	return &DaemonServer{
		rootLogger: log,
	}, nil
}

type Agent struct {
	daemonServer *DaemonServer
	httpServer   *http.Server
	conf         *Config
	logger       logr.Logger
	Status       AgentStatus
}

type AgentStatus string

func BuildAgent(conf *Config, log logr.Logger) (*Agent, error) {
	agent := &Agent{conf: conf, logger: log}
	var err error
	agent.daemonServer, err = newDaemonServer(log)
	if err != nil {
		return nil, errors.Wrap(err, "create daemon agent")
	}
	agent.httpServer = &http.Server{}
	return agent, nil
}

func (a *Agent) Start() error {
	var eg errgroup.Group
	eg.Go(func() error {
		a.logger.Info("Starting http endpoint", "address", a.conf.HttpAddr())
		if err := a.httpServer.ListenAndServe(); err != nil {
			return errors.Wrap(err, "start http endpoint")
		}
		return nil
	})
	return eg.Wait()
}
