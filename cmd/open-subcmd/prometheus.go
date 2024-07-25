package open

import (
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var PrometheusCmd = &cobra.Command{
	Use:   "prometheus",
	Short: "Opens the prometheus UI in the default browser.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return prometheus()
	},
}

func prometheus() error {
	return browser.OpenURL("http://localhost:9090")
}
