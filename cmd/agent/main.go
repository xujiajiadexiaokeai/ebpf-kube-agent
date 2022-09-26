package main

import (
	"flag"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	"xujiajiadexiaokeai.github.com/ebpf-kube-agent/pkg/agent"
	"xujiajiadexiaokeai.github.com/ebpf-kube-agent/pkg/log"
)

var (
	conf = &agent.Config{Host: "0.0.0.0"}
)

func init() {
	flag.IntVar(&conf.HTTPPort, "http-port", 2333, "the port which http server listens on")
}

func main() {
	rootLogger, err := log.NewDefaultZapLogger()
	if err != nil {
		stdlog.Fatal("failed to create root logger", err)
	}

	rootLogger = rootLogger.WithName("ebpf-agent")

	agent, err := agent.BuildAgent(conf, rootLogger)
	if err != nil {
		rootLogger.Error(err, "build ebpf-agent")
		os.Exit(1)
	}

	errs := make(chan error)
	go func() {
		errs <- agent.Start()
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	select {
	case sig := <-sigc:
		rootLogger.Info("received signal", "signal", sig)
	case err = <-errs:
		if err != nil {
			rootLogger.Error(err, "ebpf-agent server")
		}
	}
}
