package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	branchName := os.Args[1]

	defaultBranchName, err := getDefaultBranchName()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting default branch name: %v\n", err)
		return
	}

	exec.Command("git", "gone").Run()

	currentBranchName, err := getCurrentBranchName()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current branch name: %v\n", err)
		return
	}

	if currentBranchName == defaultBranchName {
		exec.Command("git", "pull", "origin", defaultBranchName).Run()
	} else {
		exec.Command("git", "fetch", "origin", defaultBranchName+":"+defaultBranchName).Run()
	}

	exec.Command("git", "checkout", defaultBranchName).Run()
	exec.Command("git", "branch", "-D", "kdy1/"+branchName).Run()
	exec.Command("git", "checkout", "origin/"+defaultBranchName, "-b", "kdy1/"+branchName).Run()
	exec.Command("git", "push", "-u", "origin", "kdy1/"+branchName).Run()
}

func getDefaultBranchName() (string, error) {
	// git symbolic-ref refs/remotes/origin/HEAD
	output, err := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get default branch name: %w", err)
	}
	return strings.TrimPrefix(string(output), "refs/remotes/origin/"), nil
}

func getCurrentBranchName() (string, error) {
	output, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch name: %w", err)
	}
	return string(output), nil
}
