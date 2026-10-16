// TextGrid: a monospaced character grid for logs and terminal-like output.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TextGrid")
	w.Resize(geo.NewSize(480, 240))

	//docs:start grid
	log := widget.NewTextGridFromString("12:00:01 starting\n12:00:02 listening on :8080")
	log.ShowLineNumbers = true

	log.Append("12:00:05 ERROR connection refused")
	red := &widget.CustomTextGridStyle{FGColor: color.NRGBA{R: 0xd0, A: 0xff}}
	log.SetRowStyle(2, red)
	log.ScrollToBottom()
	//docs:end grid

	w.SetContent(log)
	bellui.ShowAndRun(w)
}
