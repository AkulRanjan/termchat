//go:build darwin

package main

import (
	"fmt"
	"os/exec"
)

func notify(title, body string) {
	bin, err := exec.LookPath("terminal-notifier")
	if err == nil {
		cmd := exec.Command(bin, "-title", title, "-message", body)
		_ = cmd.Run()
		return
	}

	bin, err = exec.LookPath("osascript")
	if err != nil {
		return
	}

	script := fmt.Sprintf("display notification %q with title %q", body, title)
	cmd := exec.Command(bin, "-e", script)
	_ = cmd.Run()
}
