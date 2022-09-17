//go:build !android
// +build !android

package mobile

import (
	"github.com/dmowcomber/fyne/v2"
	"github.com/dmowcomber/fyne/v2/storage"
)

func nativeURI(path string) fyne.URI {
	uri, err := storage.ParseURI(path)
	if err != nil {
		fyne.LogError("Error on parsing uri", err)
	}
	return uri
}
