package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func installTools() {
	installIterm2()
	installAltTab()
	installVscode()
	installDuti()
	installOhMyZsh()
	installPodmanDesktop()
	installObsidian()
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

func installPodmanDesktop() {
	maybe("Installing Podman Desktop", exec.Command("brew", "install", "--cask", "podman-desktop"))
}

func installDuti() {
	maybe("Installing Duti", exec.Command("brew", "install", "duti"))

	// TODO: Set default editor as vscode for various file types
}

func installObsidian() {
	maybe("Installing Obsidian", exec.Command("brew", "install", "--cask", "obsidian"))
}

func installOhMyZsh() {
	omzDir := os.ExpandEnv("$HOME/.oh-my-zsh")
	if _, err := os.Stat(omzDir); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "%s already installed\n", color.GreenString("Oh My Zsh"))
		return
	}

	maybe("Installing Oh My Zsh", exec.Command("zsh", "-c", "sh -c \"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)\""))
}
