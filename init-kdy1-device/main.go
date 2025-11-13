package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	installTools()

	systemConfig()
}

func maybe(operation string, cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "%s...\n", color.GreenString(operation))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run command: %s\n%v\n", cmd.String(), err)
	}
}
