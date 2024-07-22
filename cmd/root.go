package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "otel",
	Short: "A tool that helps collect and display metric and tracing data for Spin apps.",
}

var Verbose bool
var Source string
var otelConfigDir = "spin-otel-config"

func Execute() error {
	return rootCmd.Execute()
}
