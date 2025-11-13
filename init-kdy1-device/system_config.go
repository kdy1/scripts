package main

import "os/exec"

func systemConfig() {
	disableEmojiShortcut()
}

func disableEmojiShortcut() {
	maybe("Disabling Emoji Shortcut", exec.Command("defaults", "write", "-g", "NSUserKeyEquivalents", "-dict-add", "Emoji & Symbols", "\\0"))
}
