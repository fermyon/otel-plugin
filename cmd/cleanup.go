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
	fmt.Println("Deleting Spin OTel Docker containers...")

	getContainerIDs := exec.Command("docker", "ps")
	dockerPsOutput, err := getContainerIDs.CombinedOutput()
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

	fmt.Println("Deleting Spin OTel Docker network...")

	getNetworkIDs := exec.Command("docker", "network", "list")
	networkIDOutput, err := getNetworkIDs.CombinedOutput()
	if err != nil {
		fmt.Println(string(networkIDOutput))
		return err
	}

	networkIDs := getIDs(networkIDOutput)

	// The `docker network rm` command will throw an error if there are no networks to delete
	if len(networkIDs) != 0 {
		rmNetworks := exec.Command("docker", append([]string{"network", "rm"}, networkIDs...)...)
		rmNetworksOutput, err := rmNetworks.CombinedOutput()
		if err != nil {
			fmt.Println(string(rmNetworksOutput))
			return err
		}
	}

	fmt.Println("All Spin OTel resources have been removed.")
	return nil
}
