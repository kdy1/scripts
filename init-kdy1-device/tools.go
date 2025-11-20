package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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
	installDiscord()
	installClaudeCode()
	installFnm()
	installRustup()
	installGitUtils()
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

	// Use cursor as default editor for various file types

	// osascript -e 'id of app "Cursor"'
	cursorAppId, err := getCursorAppId()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get Cursor app ID: %v\n", err)
		return
	}
	log.Println("Cursor app ID:", cursorAppId)

	handleExt := func(ext string) {
		maybe("Setting default editor for "+ext+" to Cursor", exec.Command("duti", "-s", cursorAppId, ext, "all"))
	}

	handleExt("html")
	handleExt("css")
	handleExt("js")
	handleExt("mjs")
	handleExt("cjs")
	handleExt("ts")
	handleExt("tsx")
	handleExt("json")
	handleExt("yaml")
	handleExt("yml")
	handleExt("toml")
	handleExt("md")
	handleExt("txt")

	handleExt("go")
	handleExt("rs")
	handleExt("toml")
	handleExt("md")
	handleExt("txt")
}

func getCursorAppId() (string, error) {
	cursorAppId := exec.Command("osascript", "-e", "'id of app \"Cursor\"'")
	output, err := cursorAppId.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get Cursor app ID: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func installObsidian() {
	maybe("Installing Obsidian", exec.Command("brew", "install", "--cask", "obsidian"))
}

func installDiscord() {
	maybe("Installing Discord", exec.Command("brew", "install", "--cask", "discord"))
}

func installClaudeCode() {
	maybe("Installing Claude Code", exec.Command("brew", "install", "--cask", "claude-code"))
}

func installFnm() {
	maybe("Installing Fnm", exec.Command("brew", "install", "fnm"))
}

func installRustup() {
	maybe("Installing Rustup", exec.Command("brew", "install", "rustup"))
}

func installGitUtils() {
	maybe("Installing Git", exec.Command("brew", "install", "git"))

	maybe("Installing Git Credential Manager", exec.Command("brew", "install", "--cask", "git-credential-manager"))

	maybe("Install GitHub CLI", exec.Command("brew", "install", "gh"))
}

func installOhMyZsh() {
	omzDir := os.ExpandEnv("$HOME/.oh-my-zsh")
	if _, err := os.Stat(omzDir); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "%s already installed\n", color.GreenString("Oh My Zsh"))
		return
	}

	maybe("Installing Oh My Zsh", exec.Command("zsh", "-c", "sh -c \"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)\""))
}
