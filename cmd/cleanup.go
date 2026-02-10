package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	removeContainers = false
	cleanUpCmd       = &cobra.Command{
		Use:   "cleanup",
		Short: "Clean up OpenTelemetry dependencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cleanUp(); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	cleanUpCmd.Flags().BoolVarP(&removeContainers, "remove", "r", false, "If specified, OTel containers will be removed")
}

func getIDs(runtimeOutput []byte) []string {
	var result []string
	outputArray := strings.Split(string(runtimeOutput), "\n")

	// skip the first line of the output, because it is the table header row
	// {{runtime}} ps does not support hiding column headers
	for _, entry := range outputArray[1:] {
		fields := strings.Fields(entry)
		if len(fields) > 0 {
			result = append(result, fields[0])
		}
	}

	return result
}

func cleanUp() error {
	runtime, err := detectContainerRuntime()
	if err != nil {
		return err
	}

	fmt.Println("Stopping Spin OpenTelemetry containers...")
	getContainers := exec.Command(runtime, "ps", fmt.Sprintf("--filter=name=%s*", otelConfigDirName))
	processOutput, err := getContainers.CombinedOutput()
	if err != nil {
		fmt.Println(string(processOutput))
		return err
	}

	containerIDs := getIDs(processOutput)

	// The `{{runtime}} stop` command will throw an error if there are no containers to stop
	if len(containerIDs) == 0 {
		fmt.Println("No Spin OpenTelemetry resources found. Nothing to clean up.")
		return nil
	}

	cleanupArgs := []string{"stop"}
	if removeContainers {
		cleanupArgs = []string{"rm", "-f"}
	}
	stopContainers := exec.Command(runtime, append(cleanupArgs, containerIDs...)...)
	stopContainersOutput, err := stopContainers.CombinedOutput()
	if err != nil {
		fmt.Println(string(stopContainersOutput))
		return err
	}

	fmt.Println("All Spin OpenTelemetry resources have been cleaned up.")
	return nil
}
