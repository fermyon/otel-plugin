package open

import (
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var GrafanaCmd = &cobra.Command{
	Use:   "grafana",
	Short: "Opens the Grafana UI in the default browser.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return grafana()
	},
}

func grafana() error {
	return browser.OpenURL("http://localhost:5050")
}
