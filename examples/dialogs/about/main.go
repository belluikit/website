// An About dialog built from the app's metadata.
package main

import (
	"net/url"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.about")
	w := a.NewWindow("About")

	//docs:start about
	site, _ := url.Parse("https://bellui.com")
	source, _ := url.Parse("https://github.com/belluikit")
	links := []*widget.Hyperlink{
		widget.NewHyperlink("Website", site),
		widget.NewHyperlink("Source", source),
	}
	content := "A small **example** app.\n\nReleased under the BSD 3-Clause license."

	inWindow := widget.NewButton("About (dialog)", func() {
		dialog.ShowAbout(content, links, a, w)
	})
	separate := widget.NewButton("About (window)", func() {
		dialog.ShowAboutWindow(content, links, a)
	})
	//docs:end about

	w.SetContent(container.NewVBox(inWindow, separate))
	bellui.ShowAndRun(w)
}
