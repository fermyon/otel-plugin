package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fermyon/otel-plugin/cmd"
)

func main() {
	if err := checkDependencies(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmd.Execute()

}

func checkDependencies() error {
	cmd := exec.Command("docker", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker appears not to be installed, so please visit their install page and try again once installed: https://www.docker.com/products/docker-desktop/")
	}

	cmd = exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("the docker daemon appears not to be running, so please start the daemon and try again")
	}

	return nil
}
