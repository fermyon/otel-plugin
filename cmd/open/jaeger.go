package open

import (
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var JaegerCmd = &cobra.Command{
	Use:   "jaeger",
	Short: "Opens the Jaeger UI in the default browser.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return jaeger()
	},
}

func jaeger() error {
	return browser.OpenURL("http://localhost:16686")
}
