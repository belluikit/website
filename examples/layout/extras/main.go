// More ways to arrange and draw a container's children: flex rows that centre
// or butt their children together, a fixed aspect ratio, rounded clipping,
// translucency, and swapping pages by hand.
package main

import (
	"fmt"
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/layout"
	"github.com/belluikit/bellui/widget"
)

// swatch is a rectangle of a fixed minimum size.
func swatch(c color.Color, w, h float32) *canvas.Rectangle {
	r := canvas.NewRectangle(c)
	r.SetMinSize(geo.NewSize(w, h))
	return r
}

var (
	blue   = color.NRGBA{R: 0x3b, G: 0x6e, B: 0xc8, A: 0xff}
	orange = color.NRGBA{R: 0xe8, G: 0x8a, B: 0x2c, A: 0xff}
)

func flex() core.CanvasObject {
	//docs:start flex
	// A short child next to a tall one, centred across the row instead of
	// stretched to its height. ExactPadding makes a Padding of 0 mean no gap.
	row := layout.NewFlexHBoxLayout([]float32{0, 1, 0})
	row.CrossAlign = layout.CrossCenter
	row.ExactPadding = true
	row.Padding = 0
	return container.New(row,
		swatch(blue, 40, 40), swatch(orange, 40, 80), swatch(blue, 40, 40))
	//docs:end flex
}

func aspect() core.CanvasObject {
	//docs:start aspect
	// The largest 16:9 box that fits, centred. Every child gets that box.
	screen := canvas.NewRectangle(color.NRGBA{R: 0x20, G: 0x20, B: 0x28, A: 0xff})
	return container.New(layout.NewAspectStackLayout(16.0/9.0),
		screen, container.NewCenter(widget.NewLabel("16:9")))
	//docs:end aspect
}

func clip() core.CanvasObject {
	//docs:start clip
	// Children are cut to a rounded rectangle. ClipCornerRadius only works
	// together with ClipToBounds.
	photo := canvas.NewVerticalGradient(blue, orange)
	frame := container.NewStack(photo, container.NewCenter(widget.NewLabel("Clipped")))
	frame.ClipToBounds = true
	frame.ClipCornerRadius = 16
	return frame
	//docs:end clip
}

func opacity() core.CanvasObject {
	//docs:start opacity
	// Everything inside is drawn at 40% opacity.
	faded := container.NewOpacity(
		container.NewVBox(widget.NewLabel("Faded"), widget.NewButton("Button", func() {})),
		0.4)
	show := widget.NewSlider(0, 1)
	show.Step = 0.1
	show.Value = 0.4
	show.OnChanged = func(v float64) { faded.SetOpacity(float32(v)) }
	return container.NewBorder(nil, show, nil, nil, faded)
	//docs:end opacity
}

func pages() core.CanvasObject {
	//docs:start release
	// A page area whose content is swapped by hand, not with SetContent.
	pageArea := container.NewStack(widget.NewLabel("Page 1"))
	page := 1
	next := widget.NewButton("Next page", func() {
		old := pageArea.Objects[0]
		page++
		pageArea.Objects = []core.CanvasObject{widget.NewLabel(fmt.Sprintf("Page %d", page))}
		pageArea.Refresh()
		// The old page will not be shown again: free what the toolkit
		// keeps for it at the next frame.
		bellui.Release(old)
	})
	//docs:end release
	return container.NewBorder(nil, next, nil, nil, pageArea)
}

func main() {
	a := app.NewWithID("com.example.layoutextras")
	w := a.NewWindow("More layout")

	cell := func(title string, content core.CanvasObject) core.CanvasObject {
		heading := widget.NewLabelWithStyle(title, bellui.TextAlignLeading, bellui.TextStyle{Bold: true})
		return container.NewBorder(heading, nil, nil, nil, container.NewPadded(content))
	}

	w.SetContent(container.NewGridWithColumns(3,
		cell("Flex, CrossCenter", flex()),
		cell("AspectStack", aspect()),
		cell("ClipCornerRadius", clip()),
		cell("Opacity", opacity()),
		cell("Release", pages()),
	))
	w.Resize(geo.NewSize(900, 500))
	bellui.ShowAndRun(w)
}
