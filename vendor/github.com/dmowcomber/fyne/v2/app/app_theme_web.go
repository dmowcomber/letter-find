//go:build !ci && !js && !wasm && test_web_driver
// +build !ci,!js,!wasm,test_web_driver

package app

import (
	"github.com/dmowcomber/fyne/v2"
	"github.com/dmowcomber/fyne/v2/theme"
)

func defaultVariant() fyne.ThemeVariant {
	return theme.VariantDark
}
