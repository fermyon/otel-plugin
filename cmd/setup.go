package cmd

import (
	"fmt"
	"os"
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
	if err := checkDocker(); err != nil {
		return err
	}

	composeFile := path.Join(otelConfigPath, "compose.yaml")
	if _, err := os.Stat(composeFile); os.IsNotExist(err) {
		return fmt.Errorf("The \"otel-config\" directory is missing the \"compose.yaml\" file, so please consider removing and re-installing the otel plugin")
	}

	cmd := exec.Command("docker", "compose", "-f", composeFile, "up", "-d")

	fmt.Println("Pulling and running Spin OTel resources...")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}

	fmt.Println("The Spin OTel resources are now running. Be sure to run the \"spin otel cleanup\" command when you are finished using them.")
	return nil
}
