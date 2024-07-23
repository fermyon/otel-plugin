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

	for _, entry := range outputArray {
		// Each container in the Docker Compose stack will have the name of the directory
		if strings.Contains(entry, otelConfigDir) {
			fields := strings.Fields(entry)
			if len(fields) > 0 {
				result = append(result, fields[0])
			}
		}
	}

	return result
}

func cleanUp() error {
	fmt.Println("Stopping Spin OTel Docker containers...")

	getContainers := exec.Command("docker", "ps")
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
