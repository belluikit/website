// Reads the app metadata that comes from manifest.pxf.
package main

import (
	"strconv"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	//docs:start meta
	// No ID here: app.New takes it from manifest.pxf.
	a := app.New()
	meta := a.Metadata()

	info := widget.NewLabel(meta.Name + " " + meta.Version +
		" (build " + strconv.Itoa(meta.Build) + ")\n" +
		"ID: " + meta.ID + "\n" +
		"Release build: " + strconv.FormatBool(meta.Release) + "\n" +
		"Server: " + meta.Custom["ServerURL"])
	//docs:end meta

	w := a.NewWindow(meta.Name)
	w.SetContent(container.NewVBox(info))
	bellui.ShowAndRun(w)
}
