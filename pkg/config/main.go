package config

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type GitHubBillingExporterConfig struct {
	listenAddress      *string
	metricsPath        *string
	githubToken        *string
	githubOrgs         *string
	disabledCollectors *string
	logLevel           *string
	logFormat          *string
	logOutput          *string
}

func NewGitHubBillingExporterConfig() GitHubBillingExporterConfig {
	config := GitHubBillingExporterConfig{
		listenAddress: kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").
			Default(":9776").
			Envar("GBE_WEB_LISTEN_ADDRESS").
			String(),
		metricsPath: kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").
			Default("/metrics").
			Envar("GBE_WEB_TELEMETRY_PATH").
			String(),
		githubToken: kingpin.Flag("github.token", "Access token for GitHub").
			Default("").
			Envar("GBE_GITHUB_TOKEN").
			String(),
		githubOrgs: kingpin.Flag("github.orgs", "Space separated list of GitHub Organizations").
			Default("").
			Envar("GBE_GITHUB_ORGS").
			String(),
		disabledCollectors: kingpin.Flag("disabled.collectors", "Space separated list of Disabled Collectors").
			Default("").
			Envar("GBE_DISABLED_COLLECTORS").
			String(),
		logLevel: kingpin.Flag("log.level", "Sets the log level. Valid levels are debug, info, warn, error").
			Default("info").
			Envar("GBE_LOG_LEVEL").
			String(),
		logFormat: kingpin.Flag("log.format", "Sets the log format. Valid formats are json and logfmt").
			Default("logfmt").
			Envar("GBE_LOG_FORMAT").
			String(),
		logOutput: kingpin.Flag("log.output", "Sets the log output. Valid outputs are stdout and stderr").
			Default("stdout").
			Envar("GBE_LOG_OUTPUT").
			String(),
	}
	kingpin.Version("0.0.1")
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	return config
}

func (cfg GitHubBillingExporterConfig) GetListeningAccess() string {
	return *cfg.listenAddress
}

func (cfg GitHubBillingExporterConfig) GetMetricsPath() string {
	return *cfg.metricsPath
}
