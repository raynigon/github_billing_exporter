package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/raynigon/github_billing_exporter/v2/collector"
	"github.com/raynigon/github_billing_exporter/v2/pkg/config"
	"github.com/raynigon/github_billing_exporter/v2/pkg/web"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/version"
)

var (
	gracefulStop = make(chan os.Signal, 100)
)

func registerSignals() {
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGHUP)
	signal.Notify(gracefulStop, syscall.SIGQUIT)
}

func registerCollectors(logger log.Logger, config collector.CollectorConfig) {
	collector, err := collector.NewGitHubBillingCollector(config)
	if err != nil {
		_ = level.Error(logger).Log(
			"msg", "failed to create collector",
			"err", err,
		)
		os.Exit(1)
	}

	prometheus.MustRegister(collector)
	prometheus.MustRegister(version.NewCollector("github_billing_exporter"))
}

func waitForTermination(logger log.Logger) {
	level.Info(logger).Log("msg", "listening and wait for graceful stop")
	sig := <-gracefulStop
	level.Info(logger).Log("msg", "Shutting exporter", "caught sig", sig)
	time.Sleep(2 * time.Second)
	os.Exit(0)
}

func main() {
	registerSignals()
	config := config.NewGitHubBillingExporterConfig()
	logger := config.GetLogger()
	level.Info(logger).Log("msg", "Create collectors")
	registerCollectors(logger, config.GetCollectorConfig())
	level.Info(logger).Log("msg", "Starting exporter", "version", version.Info())
	level.Info(logger).Log("Build context", version.BuildContext())

	// listener for the termination signals from the OS
	go waitForTermination(logger)
	web.RunWebserver(config)
}
