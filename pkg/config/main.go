package config

import (
	"github.com/alecthomas/kong"
)

type GitHubBillingExporterConfig struct {
	ListenAddress      string `name:"web.listen-address" default:":9776" env:"GBE_WEB_LISTEN_ADDRESS" help:"Address to listen on for web interface and telemetry."`
	MetricsPath        string `name:"web.telemetry-path" default:"/metrics" env:"GBE_WEB_TELEMETRY_PATH" help:"Path under which to expose metrics."`
	GithubToken        string `name:"github.token" default:"" env:"GBE_GITHUB_TOKEN" help:"Access token for GitHub"`
	GithubOrgs         string `name:"github.orgs" default:"" env:"GBE_GITHUB_ORGS" help:"Space separated list of GitHub Organizations"`
	DisabledCollectors string `name:"disabled.collectors" default:"" env:"GBE_DISABLED_COLLECTORS" help:"Space separated list of Disabled Collectors"`
	LogLevel           string `name:"log.level" default:"info" env:"GBE_LOG_LEVEL" help:"Sets the log level. Valid levels are debug, info, warn, error"`
	LogFormat          string `name:"log.format" default:"logfmt" env:"GBE_LOG_FORMAT" help:"Sets the log format. Valid formats are json and logfmt"`
	LogOutput          string `name:"log.output" default:"stdout" env:"GBE_LOG_OUTPUT" help:"Sets the log output. Valid outputs are stdout and stderr"`
}

func NewGitHubBillingExporterConfig() GitHubBillingExporterConfig {
	config := GitHubBillingExporterConfig{}
	kong.Parse(&config,
		kong.Name("github_billing_exporter"),
		kong.UsageOnError(),
	)
	return config
}

func (cfg GitHubBillingExporterConfig) GetListeningAccess() string {
	return cfg.ListenAddress
}

func (cfg GitHubBillingExporterConfig) GetMetricsPath() string {
	return cfg.MetricsPath
}
