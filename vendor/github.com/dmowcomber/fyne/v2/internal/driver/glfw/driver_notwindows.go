//go:build !windows
// +build !windows

package glfw

import "github.com/dmowcomber/fyne/v2"

func logError(msg string, err error) {
	fyne.LogError(msg, err)
}
