// The colour picker dialog, in its simple and advanced forms.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Colour picker")

	swatch := canvas.NewRectangle(color.NRGBA{R: 0x33, G: 0x66, B: 0xcc, A: 0xff})
	swatch.SetMinSize(geo.NewSize(120, 60))
	setColour := func(c color.Color) {
		swatch.FillColor = c
		swatch.Refresh()
	}

	//docs:start simple
	simple := widget.NewButton("Pick a colour", func() {
		dialog.ShowColorPicker("Colour", "Choose a fill colour", setColour, w)
	})
	//docs:end simple

	//docs:start advanced
	advanced := widget.NewButton("Pick any colour", func() {
		d := dialog.NewColorPicker("Colour", "Choose a fill colour", setColour, w)
		d.Advanced = true
		d.SetColor(swatch.FillColor) // start from the current colour
		d.Show()
	})
	//docs:end advanced

	w.SetContent(container.NewVBox(simple, advanced, swatch))
	bellui.ShowAndRun(w)
}
