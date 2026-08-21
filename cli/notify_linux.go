//go:build linux

package main

import (
	"os/exec"
)

func notify(title, body string) {
	bin, err := exec.LookPath("notify-send")
	if err != nil {
		return
	}

	cmd := exec.Command(bin, title, body)
	_ = cmd.Run()
}
