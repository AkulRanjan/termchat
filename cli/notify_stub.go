//go:build !linux && !darwin && !windows

package main

func notify(title, body string) {
}
