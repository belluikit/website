// Uses the resources that "bellui bundle" generated in bundled.go.
package main

//go:generate bellui bundle -o bundled.go logo.svg
//go:generate bellui bundle -o bundled.go -a welcome.txt

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Bundled resources")

	//docs:start use
	a.SetIcon(resourceLogoSvg) // a bellui.Resource, like any other
	logo := widget.NewIcon(resourceLogoSvg)
	text := widget.NewLabel(string(resourceWelcomeTxt.Content()))
	//docs:end use

	w.SetContent(container.NewVBox(logo, text))
	bellui.ShowAndRun(w)
}
