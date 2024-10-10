package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var cleanUpCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up OTel dependencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cleanUp(); err != nil {
			return err
		}
		return nil
	},
}

func getIDs(dockerOutput []byte) []string {
	var result []string
	outputArray := strings.Split(string(dockerOutput), "\n")

	// skip the first line of the output, because it is the table header row
	// docker ps does not support hiding column headers
	for _, entry := range outputArray[1:] {
		fields := strings.Fields(entry)
		if len(fields) > 0 {
			result = append(result, fields[0])
		}
	}

	return result
}

func cleanUp() error {
	if err := checkDocker(); err != nil {
		return err
	}

	fmt.Println("Stopping Spin OTel Docker containers...")
	getContainers := exec.Command("docker", "ps", fmt.Sprintf("--filter=name=%s*", otelConfigDirName), "--format=table")
	dockerPsOutput, err := getContainers.CombinedOutput()
	if err != nil {
		fmt.Println(string(dockerPsOutput))
		return err
	}

	containerIDs := getIDs(dockerPsOutput)

	// The `docker stop` command will throw an error if there are no containers to stop
	if len(containerIDs) != 0 {
		stopContainers := exec.Command("docker", append([]string{"stop"}, containerIDs...)...)
		stopContainersOutput, err := stopContainers.CombinedOutput()
		if err != nil {
			fmt.Println(string(stopContainersOutput))
			return err
		}
	}

	fmt.Println("All Spin OTel resources have been cleaned up.")
	return nil
}
