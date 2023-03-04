package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/xujiajiadexiaokeai/ebpf-kube-agent/pkg/manager"
)

var (
	conf = &manager.Config{Host: "0.0.0.0"}
)

func init() {
	flag.IntVar(&conf.GrpcPort, "gprc-port", 3311, "gprc port")

	flag.Parse()
}

func main() {
	logger := logrus.New()
	manager, err := manager.BuildManager(conf, *logger)
	if err != nil {
		logger.Error(err, "build manager failed")
		os.Exit(1)
	}

	errs := make(chan error)
	go func() {
		errs <- manager.Start()
	}()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	select {
	case sig := <-sigc:
		logger.Info("received signal: ", "signal", sig)
	case err = <-errs:
		if err != nil {
			logger.Error(err, "ebpf-kube-manager")
		}
	}
	if err = manager.Stop(); err != nil {
		logger.Error(err, "manager stop")
	}
}
