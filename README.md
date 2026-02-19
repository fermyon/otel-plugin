# Spin OTel Plugin

This is a plugin that makes it easy to use OTel with Spin.

## Background

Spin applications have the ability to export metrics, logs, and trace data. This plugin provides dashboards for viewing the data.

## Requirements

This plugin relies on third-party software to work properly. Please be sure you have the following installed before continuing:

- Latest version of [Docker](https://www.docker.com/products/docker-desktop) or [Podman](https://podman.io/docs)

# Installation

You can install the `otel` plugin either using a stable release or from the `main` branch.

## Install the latest version of the otel plugin

The latest stable release of the `otel` plugin can be installed as shown here:

```sh
spin plugins update
spin plugin install otel
```

## Install the canary version of the plugin

The canary release of the `otel` plugin represents the most recent commit on `main` (`HEAD`) and may not be stable, with some features still in progress.

```sh
spin plugins install --url https://github.com/fermyon/otel-plugin/releases/download/canary/otel.json
```

## Install from a local build

Alternatively, use the `spin pluginify` plugin to install from a fresh build. This will use the pluginify manifest (`spin-pluginify.toml`) to package the plugin and proceed to install it:

```sh
spin plugins install pluginify
go build -o otel
spin pluginify --install
```

# Observability Stacks

The `otel` plugin currently supports two different observability stacks:

- Default: Multi-container observability stack based on Prometheus, Loki, Grafana and Jaeger
- Aspire: Single-container observability stack using .NET Aspire Dashboard

# Usage

Once the plugin is installed, you can try the below commands:

## Set up the observability stack

You chose the desired observability stack as part of the `setup` command. You can change the observability stack at any point in time by cleaning it up and re-running the `setup` command using another stack.

### Set up the default observability stack

```sh
spin otel setup
```

### Set up the aspire observability stack

```sh
spin otel setup --aspire
```


## Run a Spin app and exports telemetry data

```sh
spin otel up
```

Any flags that work with the `spin up` command, will work with the `spin otel up` command.

```sh
spin otel up -- --help
```

Additionally, the `spin otel up` command enables the `wasi-otel` features in versions of Spin >= v3.6.0, which means that users will be able to instrument their applications with OpenTelemetry using the [opentelemetry-wasi](github.com/bytecodealliance/opentelemetry-wasi) SDKs.

## Open the dashboards in the default browser

Depending on the chosen observability stack, you can use the sub-commands of `spin otel open` to open corresponding dashboards using your default browser.

Dashboard for viewing metrics and logs:

```sh
spin otel open grafana
```

Dashboard for viewing trace data:

```sh
spin otel open jaeger
```

Dashboard for querying and viewing metrics:

```sh
spin otel open prometheus
```

.NET Aspire Dashboard (all-in-one).

```sh
spin otel open aspire
```

## Terminate the dashboards

```sh
spin otel cleanup
```
