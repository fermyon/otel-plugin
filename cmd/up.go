package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Runs a Spin App with the default OTel environment variables.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := up(); err != nil {
			return err
		}
		return nil
	},
}

func up() error {
	pathToSpin := os.Getenv("SPIN_BIN_PATH")
	if pathToSpin == "" {
		return fmt.Errorf("please ensure that you are running 'spin otel up', rather than calling the OTel plugin binary directly")
	}

	cmd := exec.Command(pathToSpin, "up")
	cmd.Env = append(
		os.Environ(),
		"OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
