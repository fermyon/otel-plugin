package cmd

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the desired OTel UI in the default browser. Accepts the following arguments: 'prometheus' or 'p', 'grafana' or 'g', 'jaeger' or 'j'.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must specify the UI you wish to open: 'prometheus', 'grafana', or 'jaeger'")
		}

		if err := open(args[0]); err != nil {
			return err
		}
		return nil
	},
}

func open(param string) error {
	var url string
	if param == "p" || param == "prometheus" {
		url = "http://localhost:9090"
	} else if param == "g" || param == "grafana" {
		url = "http://localhost:5050"
	} else if param == "j" || param == "jaeger" {
		url = "http://localhost:16686"
	}

	if err := browser.OpenURL(url); err != nil {
		return err
	}

	return nil
}
