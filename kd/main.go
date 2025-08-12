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

	mustSuccess(exec.Command("git", "gone"))

	currentBranchName, err := getCurrentBranchName()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current branch name: %v\n", err)
		return
	}

	// Stash all changes
	mustSuccess(exec.Command("git", "stash", "save", "kd-stash"))

	if currentBranchName == defaultBranchName {
		mustSuccess(exec.Command("git", "pull"))
	} else {
		mustSuccess(exec.Command("git", "fetch", "origin", defaultBranchName+":"+defaultBranchName))
	}

	ignoreError(exec.Command("git", "checkout", defaultBranchName))
	ignoreError(exec.Command("git", "branch", "-D", "kdy1/"+branchName))

	mustSuccess(exec.Command("git", "checkout", "origin/"+defaultBranchName, "-b", "kdy1/"+branchName))
	mustSuccess(exec.Command("git", "push", "-u", "origin", "kdy1/"+branchName))

	// Pop the stash
	mustSuccess(exec.Command("git", "stash", "pop"))
}

func ignoreError(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running command: %s\n%v\n", cmd.String(), err)
	}
}

func mustSuccess(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running command: %s\n%v\n", cmd.String(), err)
		os.Exit(1)
	}
}

func getDefaultBranchName() (string, error) {
	// git symbolic-ref refs/remotes/origin/HEAD
	output, err := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get default branch name: %w", err)
	}
	return strings.TrimSpace(strings.TrimPrefix(string(output), "refs/remotes/origin/")), nil
}

func getCurrentBranchName() (string, error) {
	output, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch name: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}
