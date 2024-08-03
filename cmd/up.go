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
	Long:  "Runs a Spin App with the default OTel environment variables. Any flags that work with the `spin up` command, will work with the `spin otel up` command: 'spin otel up -- --help'",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := up(args); err != nil {
			return err
		}
		return nil
	},
	Args: cobra.ArbitraryArgs,
}

func up(args []string) error {
	pathToSpin := os.Getenv("SPIN_BIN_PATH")
	if pathToSpin == "" {
		return fmt.Errorf("please ensure that you are running 'spin otel up', rather than calling the OTel plugin binary directly")
	}

	// Passing flags and args after the '--'
	cmdArgs := append([]string{"up"}, args...)
	cmd := exec.Command(pathToSpin, cmdArgs...)
	cmd.Env = append(
		os.Environ(),
		"OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
