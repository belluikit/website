// The canvas primitives side by side: text, shapes, a line, gradients, a
// raster and an image, next to an ordinary widget.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.canvas")
	w := a.NewWindow("Canvas objects")

	blue := color.NRGBA{R: 0x3b, G: 0x6e, B: 0xc8, A: 0xff}
	orange := color.NRGBA{R: 0xe8, G: 0x8a, B: 0x2c, A: 0xff}

	//docs:start shapes
	text := canvas.NewText("canvas.Text", theme.Color(theme.ColorNameForeground))
	text.TextSize = 20
	text.TextStyle = bellui.TextStyle{Bold: true}

	rect := canvas.NewRectangle(blue)
	rect.CornerRadius = 12
	rect.StrokeColor = orange
	rect.StrokeWidth = 3

	circle := canvas.NewCircle(orange)

	line := canvas.NewLine(blue)
	line.StrokeWidth = 4

	gradient := canvas.NewHorizontalGradient(blue, orange)
	radial := canvas.NewRadialGradient(orange, color.Transparent)
	//docs:end shapes

	//docs:start raster
	// A raster asks for a colour per pixel, at whatever size it is drawn.
	checker := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		if (x/16+y/16)%2 == 0 {
			return color.NRGBA{R: 0xdd, G: 0xdd, B: 0xdd, A: 0xff}
		}
		return color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	})
	//docs:end raster

	//docs:start image
	img := canvas.NewImageFromResource(theme.HomeIcon())
	img.FillMode = canvas.ImageFillContain
	//docs:end image

	//docs:start widget
	// A widget is a canvas object too, so it goes anywhere the others go.
	button := widget.NewButton("widget.Button", func() {})
	//docs:end widget

	cell := func(o core.CanvasObject) core.CanvasObject {
		// Canvas objects report a minimum size of 1x1 unless told otherwise;
		// a fixed-size cell keeps each one visible.
		return container.NewGridWrap(geo.NewSize(140, 90), o)
	}
	w.SetContent(container.NewGridWithColumns(3,
		cell(text), cell(rect), cell(circle),
		cell(line), cell(gradient), cell(radial),
		cell(checker), cell(img), cell(container.NewCenter(button)),
	))
	bellui.ShowAndRun(w)
}
