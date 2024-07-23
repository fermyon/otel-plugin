package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otel",
	Short: "A plugin that makes using Spin with OTel easy.",
	Long:  "A plugin that makes using Spin with OTel easy by automatically standing up dependencies, sourcing environment variables, and linking to dashboards",
}

var otelConfigDir = "otel-config"

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(cleanUpCmd)
	rootCmd.AddCommand(setUpCmd)
	rootCmd.AddCommand(openCmd)
}
