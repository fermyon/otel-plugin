package cmd

import (
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the desired OpenTelemetry UI in the default browser.",
}
