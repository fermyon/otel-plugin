# Spin OTel Plugin

This is a plugin that makes it easy to use OTel with Spin.

## Background

Spin applications have the ability to export metrics and trace data. This plugin provides dashboards for viewing the data.

# Installation

The trigger is installed as a Spin plugin. It can be installed from a release or build.

## Install the latest version of the plugin

The latest stable release of the command trigger plugin can be installed like so:

```sh
spin plugins update
spin plugin install otel
```

## Install the canary version of the plugin

The canary release of the command trigger plugin represents the most recent commits on `main` and may not be stable, with some features still in progress.

```sh
spin plugins install --url https://github.com/fermyon/otel-plugin/releases/download/canary/otel.json
```

## Install from a local build

Alternatively, use the `spin pluginify` plugin to install from a fresh build. This will use the pluginify manifest (`spin-pluginify.toml`) to package the plugin and proceed to install it:

```sh
spin plugins install pluginify
go build -o otel
spin pluginify install
```

# Usage

Once the plugin is installed, you can try the below commands:

## Set up the dashboards

```sh
spin otel setup
```

## Run a Spin app that exports telemetry data

```sh
spin otel up
```

Any flags that work with the `spin up` command, will work with the `spin otel up` command.

```sh
spin otel up -- --help
```

## Open the dashboards in the default browser

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

## Terminate the dashboards

```sh
spin otel cleanup
```
