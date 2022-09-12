# GitHub billing exporter

[![GitHub Release](https://img.shields.io/github/release/raynigon/github_billing_exporter.svg?style=flat)](https://github.com/raynigon/github_billing_exporter/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/raynigon/github_billing_exporter/v2)](https://goreportcard.com/report/github.com/raynigon/github_billing_exporter/v2)

Forked From: https://github.com/borisputerka/github_billing_exporter because its not maintained there anymore.

This exporter exposes [Prometheus](https://prometheus.io/) metrics from GitHub billing API [endpoint](https://docs.github.com/en/free-pro-team@latest/rest/reference/billing) and the GitHub timing API [endpoint](https://docs.github.com/en/rest/reference/actions#get-workflow-usage).


## Metrics
```
# HELP github_billing_actions_org_minutes_billed_count GitHub actions used minutes with platform multipliers
# TYPE github_billing_actions_org_minutes_billed_count counter
github_billing_actions_org_minutes_billed_count{org="<PLACEHOLDER>",platform="linux"} 12345
github_billing_actions_org_minutes_billed_count{org="<PLACEHOLDER>",platform="macos"} 12345
github_billing_actions_org_minutes_billed_count{org="<PLACEHOLDER>",platform="windows"} 12345
# HELP github_billing_actions_org_minutes_inclusive GitHub actions inclusive budget minutes
# TYPE github_billing_actions_org_minutes_inclusive gauge
github_billing_actions_org_minutes_inclusive{org="<PLACEHOLDER>"} 5000
# HELP github_billing_actions_org_minutes_paid_count Total GitHub actions minutes paid for
# TYPE github_billing_actions_org_minutes_paid_count counter
github_billing_actions_org_minutes_paid_count{org="<PLACEHOLDER>"} 12345
# HELP github_billing_actions_org_minutes_real_count GitHub actions used minutes without platform multiplier
# TYPE github_billing_actions_org_minutes_real_count counter
github_billing_actions_org_minutes_real_count{org="<PLACEHOLDER>",platform="linux"} 12345
github_billing_actions_org_minutes_real_count{org="<PLACEHOLDER>",platform="macos"} 12345
github_billing_actions_org_minutes_real_count{org="<PLACEHOLDER>",platform="windows"} 12345
# HELP github_billing_actions_org_minutes_total_count Total GitHub actions minutes used
# TYPE github_billing_actions_org_minutes_total_count counter
github_billing_actions_org_minutes_total_count{org="<PLACEHOLDER>"} 12345
# HELP github_billing_actions_workflow_minutes_billed_count GitHub actions used minutes with platform multipliers
# TYPE github_billing_actions_workflow_minutes_billed_count counter
github_billing_actions_workflow_minutes_billed_count{org="<PLACEHOLDER>",platform="linux",repository="<PLACEHOLDER>",workflow_id="<PLACEHOLDER>",workflow_name="<PLACEHOLDER>"} 12345
# HELP github_billing_actions_workflow_minutes_real_count GitHub actions used minutes without platform multiplier
# TYPE github_billing_actions_workflow_minutes_real_count counter
github_billing_actions_workflow_minutes_real_count{org="<PLACEHOLDER>",platform="linux",repository="<PLACEHOLDER>",workflow_id="<PLACEHOLDER>",workflow_name="<PLACEHOLDER>"} 12345
# HELP github_billing_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which github_billing_exporter was built.
# TYPE github_billing_exporter_build_info gauge
github_billing_exporter_build_info{branch="",goversion="go1.17.13",revision="",version=""} 1
# HELP github_billing_packages_org_bandwith_inclusive GitHub packages inclusive budget bandwith in gigabytes
# TYPE github_billing_packages_org_bandwith_inclusive gauge
github_billing_packages_org_bandwith_inclusive{org="<PLACEHOLDER>"} 1234
# HELP github_billing_packages_org_bandwith_paid_count GitHub packages paid used bandwith in gigabytes
# TYPE github_billing_packages_org_bandwith_paid_count counter
github_billing_packages_org_bandwith_paid_count{org="<PLACEHOLDER>"} 1234
# HELP github_billing_packages_org_bandwith_total_count GitHub packages total used bandwith in gigabytes
# TYPE github_billing_packages_org_bandwith_total_count counter
github_billing_packages_org_bandwith_total_count{org="<PLACEHOLDER>"} 1234
# HELP github_billing_storage_org_billing_cycle_days Days left in the current billing cycle
# TYPE github_billing_storage_org_billing_cycle_days gauge
github_billing_storage_org_billing_cycle_days{org="<PLACEHOLDER>"} 1234
# HELP github_billing_storage_org_paid_count GitHub storage used paid in gigabytes
# TYPE github_billing_storage_org_paid_count counter
github_billing_storage_org_paid_count{org="<PLACEHOLDER>"} 123.45
# HELP github_billing_storage_org_total_count GitHub storage used total in gigabytes
# TYPE github_billing_storage_org_total_count counter
github_billing_storage_org_total_count{org="<PLACEHOLDER>"} 1234
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 15631
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

## Installation

For pre-built binaries please take a look at the releases.
https://github.com/raynigon/github_billing_exporter/releases

### Docker

```bash
docker pull ghcr.io/raynigon/github-billing-exporter:latest
docker run --rm -p 9776:9776 ghcr.io/raynigon/github-billing-exporter:latest
```

Example `docker-compose.yml`:

```yaml
github_billing_exporter:
    image: ghcr.io/raynigon/github-billing-exporter:latest
    command:
     - '--githun.token=<SECRET>'
    restart: always
    ports:
    - "127.0.0.1:9776:9776"
```

### Kubernetes

You can find an deployment definition at: https://github.com/raynigon/github_billing_exporter/tree/main/examples/kubernetes/deployment.yaml .

## Building and running

### Build

    make

### Running

Running using an environment variable:

    export GBE_GITHUB_ORGS="ORG1 ORG2 ..."
    export GBE_GITHUB_TOKEN="example_token"
    ./github_billing_exporter

Running using args:

    ./github_billing_exporter \
    --github-orgs="ORG1 ORG2 ..." \
    --github-token="example_token"

## Collectors

There are three collectors (`actions`, `packages` and `storage`) all enabled by default. Disabling collector(s) can be done using arg `--no-collector.<name>`.

### List of collectors

Name	          | Description									                                        | Enabled
------------------|-------------------------------------------------------------------------------------|--------
actions_org       | Exposes billing statistics from `/orgs/{org}/settings/billing/actions`	            | `true`
packages_org      | Exposes billing statistics from `/orgs/{org}/settings/billing/packages`	            | `true`
storage_org       | Exposes billing statistics from `/orgs/{org}/settings/billing/shared-storage`       | `true`
actions_workflow  | Exposes used time from `/repos/{org}/{repo}/actions/workflows/{workflow_id}/timing` | `true`

## Environment variables / args reference

Version    | Env		               | Arg		             | Description			                	       | Default
-----------|---------------------------|-------------------------|-------------------------------------------------|---------
\>=`0.3.0` | `GBE_LISTEN_ADDRESS`      | `--web.listen-address`  | Address on which to expose metrics.             | `:9776`
\>=`0.3.0` | `GBE_METRICS_PATH`	       | `--web.telemetry-path`  | Path under which to expose metrics.             | `/metrics`
\>=`0.3.0` | `GBE_GITHUB_TOKEN`        | `--github.token`	     | GitHub token with billing/repo read privileges  | `""`
\>=`0.3.0` | `GBE_GITHUB_ORGS`	       | `--github.orgs`	     | GitHub organizations to scrape metrics for      | `""`
\>=`0.3.0` | `GBE_LOG_LEVEL`           | `--log.level`	         | -                                               | `"info"`
\>=`0.3.0` | `GBE_LOG_FORMAT`          | `--log.format`	         | -                                               | `"logfmt"`
\>=`0.3.0` | `GBE_LOG_OUTPUT`          | `--log.output`	         | -                                               | `"stdout"`
\>=`0.4.0` | `GBE_DISABLED_COLLECTORS` | `--disabled-collectors` | Collectors to disable			               | `""`

### Token privileges

The GitHub Private Access Token needs to have access to read billing data, repositories and workflows.