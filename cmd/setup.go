package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/fermyon/otel-plugin/internal/stack"
	"github.com/spf13/cobra"
)

var (
	aspire   = false
	setUpCmd = &cobra.Command{
		Use:   "setup",
		Short: "Run OpenTelemetry dependencies as containers",
		Long:  "Required OpenTelemetry dependencies will be started as containers using a supported container runtime (docker, podman)",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := stack.GetStackByFlags(aspire)
			if err := setUp(s); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	setUpCmd.PersistentFlags().BoolVarP(&aspire, "aspire", "", false, "Use .NET Aspire dashboard as OTel stack")
}

func setUp(s stack.Stack) error {
	runtime, err := detectContainerRuntime()
	if err != nil {
		return err
	}

	composeFileName := s.GetComposeFileName()
	composeFilePath := path.Join(otelConfigPath, composeFileName)
	if _, err := os.Stat(composeFilePath); os.IsNotExist(err) {
		return fmt.Errorf("The \"otel-config\" directory is missing the \"%s\" file, so please consider removing and re-installing the otel plugin", composeFileName)
	}

	cmd := exec.Command(runtime, "compose", "-f", composeFilePath, "up", "-d")

	fmt.Println("Pulling and running Spin OpenTelemetry resources...")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}

	fmt.Println("The Spin OpenTelemetry resources are now running. Be sure to run the \"spin otel cleanup\" command when you are finished using them.")
	return nil
}
