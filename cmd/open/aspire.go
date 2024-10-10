package open

import (
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var AspireCmd = &cobra.Command{
	Use:   "aspire",
	Short: "Opens the .NET Aspire Dashboard in the default browser.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return aspire()
	},
}

func aspire() error {
	return browser.OpenURL("http://localhost:18888")
}
