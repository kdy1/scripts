package main

import "os/exec"

func systemConfig() {
	configureGit()
	configureMacOS()
}

func configureMacOS() {
	maybe("Disabling Emoji Shortcut", exec.Command("defaults", "write", "-g", "NSUserKeyEquivalents", "-dict-add", "Emoji & Symbols", "\\0"))
	// sudo defaults write /Library/Preferences/FeatureFlags/Domain/UIKit.plist emoji_enhancements -dict-add Enabled -bool NO
	maybe("Disabling Emoji Enhancements", exec.Command("defaults", "write", "/Library/Preferences/FeatureFlags/Domain/UIKit.plist", "emoji_enhancements", "-dict-add", "Enabled", "-bool", "NO"))

	// Enable dock autohide
	maybe("Enabling dock autohide", exec.Command("defaults", "write", "com.apple.dock", "autohide", "-bool", "TRUE"))
	// Disable recent items in dock
	maybe("Disabling recent items in dock", exec.Command("defaults", "write", "com.apple.dock", "show-recents", "-bool", "FALSE"))

	// restart dock
	maybe("Restarting dock", exec.Command("killall", "Dock"))
}

func configureGit() {
	// mkdir -p ~/projects
	maybe("Creating ~/projects directory", exec.Command("mkdir", "-p", "~/projects"))

	// git config --global user.name "DongYun Kang"
	maybe("Configuring Git user name", exec.Command("git", "config", "--global", "user.name", "DongYun Kang"))

	// git config --global user.email "kdy.1997.dev@gmail.com"
	maybe("Configuring Git user email", exec.Command("git", "config", "--global", "user.email", "kdy.1997.dev@gmail.com"))
}
