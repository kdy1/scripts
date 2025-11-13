package main

import "os/exec"

func systemConfig() {
	configureGit()
	disableEmojiShortcut()
}

func disableEmojiShortcut() {
	maybe("Disabling Emoji Shortcut", exec.Command("defaults", "write", "-g", "NSUserKeyEquivalents", "-dict-add", "Emoji & Symbols", "\\0"))
}

func configureGit() {
	// mkdir -p ~/projects
	maybe("Creating ~/projects directory", exec.Command("mkdir", "-p", "~/projects"))

	// git config --global user.name "DongYun Kang"
	maybe("Configuring Git user name", exec.Command("git", "config", "--global", "user.name", "DongYun Kang"))

	// git config --global user.email "kdy.1997.dev@gmail.com"
	maybe("Configuring Git user email", exec.Command("git", "config", "--global", "user.email", "kdy.1997.dev@gmail.com"))
}
