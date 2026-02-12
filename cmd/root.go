package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Masterminds/semver"
	open "github.com/fermyon/otel-plugin/cmd/open"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "otel",
	Short: "A plugin that makes using Spin with OpenTelemetry easy.",
	Long:  "A plugin that makes using Spin with OpenTelemetry easy by automatically standing up dependencies, sourcing environment variables, and linking to dashboards.",
}

var otelConfigDirName = "otel-config"
var otelConfigPath string

// setOtelConfigPath allows for someone to run the otel plugin directly from source or via the Spin plugin installation
func setOtelConfigPath() error {
	executablePath, err := os.Executable()
	if err != nil {
		return err
	}

	otelConfigPath = path.Join(path.Dir(executablePath), otelConfigDirName)

	if _, err := os.Stat(otelConfigPath); os.IsNotExist(err) {
		return fmt.Errorf("The directory in which the plugin binary is executed is missing necessary files, so please make sure the plugin was installed using \"spin plugins install otel\"")
	}

	return nil
}

// detectContainerRuntime checks for a supported container runtime.
//
// Returns the detected runtime name.
func detectContainerRuntime() (string, error) {
	// Default Docker, fallback Podman
	runtimesToCheck := []string{"docker", "podman"}
	runtime := ""
	for _, rt := range runtimesToCheck {
		cmd := exec.Command(rt, "--version")
		if err := cmd.Run(); err != nil {
			continue
		}
		runtime = rt
		break
	}

	if runtime == "" {
		return "", fmt.Errorf("Unable to detect container runtime.\nPlease ensure at least one of the following is installed and in the PATH: %s\n", strings.Join(runtimesToCheck, ", "))
	}

	if runtime == "podman" {
		// Podman doesn't auto-install a compose provider
		composeProvidersToCheck := []string{"podman-compose", "docker-compose"}
		composeProvider := ""
		for _, rt := range composeProvidersToCheck {
			cmd := exec.Command(rt, "--version")
			if err := cmd.Run(); err != nil {
				continue
			}

			composeProvider = rt
			break
		}

		if composeProvider == "" {
			return "", fmt.Errorf("Unable to detect compose provider.\n Please ensure at least one of the following is installed and in the PATH: %s\n", strings.Join(composeProvidersToCheck, ", "))
		}
	}

	cmd := exec.Command(runtime, "info")
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("The daemon/virtual machine for %[1]s appears not to be running. Please check the documentation for %[1]s for more information.", runtime)
	}

	return runtime, nil
}

func getSpinVersion() (*semver.Version, error) {
	spin, err := getSpinPath()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(spin, "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	outputParts := strings.Split(string(output), " ")
	if len(outputParts) < 2 {
		return nil, fmt.Errorf("Spin appears to have changed how it formats its version flag output. Expected \"spin {{VERSION}}\", got %s", string(output))
	}

	semver, err := semver.NewVersion(outputParts[1])
	if err != nil {
		return nil, fmt.Errorf("Spin version is not valid semver: %v", err)
	}

	return semver, nil
}

func getSpinPath() (string, error) {
	pathToSpin := os.Getenv("SPIN_BIN_PATH")
	if pathToSpin == "" {
		return "", fmt.Errorf("Please ensure that you are running \"spin otel up\", rather than calling the OpenTelemetry plugin binary directly")
	}

	return pathToSpin, nil
}

func Execute() {
	if err := setOtelConfigPath(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("Error finding the \"otel-config\" directory: %w", err))
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(cleanUpCmd)
	rootCmd.AddCommand(setUpCmd)
	rootCmd.AddCommand(openCmd)
	openCmd.AddCommand(open.GrafanaCmd)
	openCmd.AddCommand(open.JaegerCmd)
	openCmd.AddCommand(open.PrometheusCmd)
	openCmd.AddCommand(open.AspireCmd)
}
