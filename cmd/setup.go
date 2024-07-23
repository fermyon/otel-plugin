package cmd

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

var setUpCmd = &cobra.Command{
	Use:   "setup",
	Short: "Run OTel dependencies in Docker.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := setUp(); err != nil {
			return err
		}
		return nil
	},
}

func setUp() error {
	cmd := exec.Command("docker", "compose", "-f", path.Join(otelConfigDir, "compose.yaml"), "up", "-d")

	fmt.Println("Pulling and running Spin OTel resources...")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}

	fmt.Println("The Spin OTel resources are now running. Be sure to run the `spin otel cleanup` command when you are finished using them.")
	return nil
}
