package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	installIterm2()
	installAltTab()
	installVscode()
	installOhMyZsh()

	systemConfig()
}

func installIterm2() {
	maybe("Installing Iterm2", exec.Command("brew", "install", "--cask", "iterm2"))
}

func installAltTab() {
	maybe("Installing AltTab", exec.Command("brew", "install", "--cask", "alt-tab"))
}

func installVscode() {
	maybe("Installing VSCode", exec.Command("brew", "install", "--cask", "visual-studio-code"))
}

func installOhMyZsh() {
	omzDir := os.ExpandEnv("$HOME/.oh-my-zsh")
	if _, err := os.Stat(omzDir); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "%s already installed\n", color.GreenString("Oh My Zsh"))
		return
	}

	maybe("Installing Oh My Zsh", exec.Command("zsh", "-c", "sh -c \"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)\""))
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
