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
	flag.IntVar(&conf.GrpcPort, "gprc-port", 0, "gprc port")

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
	if err = manager.Stop(); err != nil {
		logger.Error(err, "manager stop")
	}
}
