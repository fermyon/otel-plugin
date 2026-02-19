package cmd

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var experimentalWasiOtel = false

func init() {
	upCmd.Flags().BoolVarP(&experimentalWasiOtel, "experimental-wasi-otel", "", false, "Enable the experimental wasi-otel feature in Spin")
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Runs a Spin App with the default OpenTelemetry environment variables and wasi-otel features enabled.",
	Long:  "Runs a Spin App with the default OpenTelemetry environment variables and wasi-otel features enabled. Any flags that work with the \"spin up\" command, will work with the \"spin otel up\" command: \"spin otel up -- --help\"",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := up(args); err != nil {
			return err
		}
		return nil
	},
	Args: cobra.ArbitraryArgs,
}

func up(args []string) error {
	pathToSpin, err := getSpinPath()
	if err != nil {
		return err
	}

	cmdArgs := []string{"up"}

	if spinHasWasiOtel() {
		cmdArgs = append(cmdArgs, "--experimental-wasi-otel")
	}

	// Passing flags and args after the '--'
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command(pathToSpin, cmdArgs...)
	cmd.Env = append(
		os.Environ(),
		"OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318",
		"OTEL_EXPORTER_OTLP_HEADERS=x-otlp-api-key=SpinOTelApiKey",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func spinHasWasiOtel() bool {
	major, err := strconv.Atoi(os.Getenv("SPIN_VERSION_MAJOR"))
	if err != nil {
		panic("Something went wrong parsing the SPIN_VERSION_MAJOR env var: " + err.Error())
	}

	minor, err := strconv.Atoi(os.Getenv("SPIN_VERSION_MINOR"))
	if err != nil {
		panic("Something went wrong parsing the SPIN_VERSION_MINOR env var: " + err.Error())
	}

	if major < 3 {
		return false
	} else if major > 3 {
		return true
	} else { // if major == 3
		if minor >= 6 {
			return true
		} else {
			return false
		}
	}
}
