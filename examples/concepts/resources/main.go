// Images from three sources: a file embedded in the binary, the theme's
// icon set, and a file on disk named on the command line.
package main

import (
	_ "embed"
	"os"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/storage"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

//docs:start embed
//go:embed logo.svg
var logoSVG []byte

// The name identifies the resource; keep the file extension, so an SVG
// is recognised as one.
var logo = bellui.NewStaticResource("logo.svg", logoSVG)

//docs:end embed

func main() {
	a := app.NewWithID("com.example.resources")
	w := a.NewWindow("Resources")

	//docs:start use
	// One resource, used three ways.
	a.SetIcon(logo) // the app icon, and the default icon of new windows

	embedded := canvas.NewImageFromResource(logo)
	embedded.FillMode = canvas.ImageFillContain

	button := widget.NewButtonWithIcon("Logo", logo, func() {})
	//docs:end use

	//docs:start theme
	// Theme icons follow the current theme's colours and switch with it.
	home := widget.NewIcon(theme.HomeIcon())

	// Wrap your own monochrome icon to give it the theme's primary colour.
	tinted := widget.NewIcon(theme.NewPrimaryThemedResource(logo))
	//docs:end theme

	items := []core.CanvasObject{embedded, button, home, tinted}

	//docs:start file
	if len(os.Args) > 1 {
		// A file:// URI is read straight from disk.
		fromDisk := canvas.NewImageFromURI(storage.NewFileURI(os.Args[1]))
		fromDisk.FillMode = canvas.ImageFillContain
		items = append(items, fromDisk)
	}
	//docs:end file

	w.SetContent(container.NewGridWrap(geo.NewSize(96, 96), items...))
	w.Resize(geo.NewSize(480, 240))
	bellui.ShowAndRun(w)
}
