// Every standard container in one window, each in its own titled cell.
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

// box is a coloured rectangle with a label, used as sample content.
func box(name string) core.CanvasObject {
	bg := canvas.NewRectangle(color.NRGBA{R: 0x3b, G: 0x6e, B: 0xc8, A: 0x60})
	return container.NewStack(bg, widget.NewLabel(name))
}

func vbox() core.CanvasObject {
	//docs:start vbox
	// Top to bottom. Each child is full width and only as tall as it needs.
	return container.NewVBox(box("One"), box("Two"), box("Three"))
	//docs:end vbox
}

func hbox() core.CanvasObject {
	//docs:start hbox
	// Left to right. The spacer takes the free space, pushing "Right" to the end.
	return container.NewHBox(box("Left"), layout.NewSpacer(), box("Right"))
	//docs:end hbox
}

func border() core.CanvasObject {
	//docs:start border
	// Edges first (top, bottom, left, right; nil leaves one empty),
	// then the content that fills the middle.
	return container.NewBorder(box("Top"), box("Bottom"), box("L"), nil, box("Middle"))
	//docs:end border
}

func grid() core.CanvasObject {
	//docs:start grid
	// Equal-sized cells, two per row. Rows are added as needed.
	return container.NewGridWithColumns(2, box("1"), box("2"), box("3"), box("4"))
	//docs:end grid
}

func gridWrap() core.CanvasObject {
	//docs:start gridwrap
	// Every child gets the same fixed size; they flow onto new rows
	// when the width runs out.
	var tiles []core.CanvasObject
	for i := 1; i <= 6; i++ {
		tiles = append(tiles, box(fmt.Sprint(i)))
	}
	return container.NewGridWrap(geo.NewSize(50, 40), tiles...)
	//docs:end gridwrap
}

func center() core.CanvasObject {
	//docs:start center
	// The child keeps its minimum size and sits in the middle.
	return container.NewCenter(widget.NewButton("Centred", func() {}))
	//docs:end center
}

func stack() core.CanvasObject {
	//docs:start stack
	// Children on top of each other, all at full size. Later ones are on top.
	bg := canvas.NewVerticalGradient(
		color.NRGBA{R: 0x3b, G: 0x6e, B: 0xc8, A: 0xff},
		color.NRGBA{R: 0xe8, G: 0x8a, B: 0x2c, A: 0xff})
	return container.NewStack(bg, container.NewCenter(widget.NewLabel("On top")))
	//docs:end stack
}

func padded() core.CanvasObject {
	//docs:start padded
	// The child fills the space minus the theme's padding on every side.
	return container.NewPadded(box("Padded"))
	//docs:end padded
}

func form() core.CanvasObject {
	//docs:start form
	// Two columns: labels on the left, content on the right.
	return container.New(layout.NewFormLayout(),
		widget.NewLabel("Name"), widget.NewEntry(),
		widget.NewLabel("Email"), widget.NewEntry(),
	)
	//docs:end form
}

func main() {
	a := app.NewWithID("com.example.containers")
	w := a.NewWindow("Containers")

	cell := func(title string, content core.CanvasObject) core.CanvasObject {
		heading := widget.NewLabelWithStyle(title, bellui.TextAlignLeading, bellui.TextStyle{Bold: true})
		frame := canvas.NewRectangle(color.Transparent)
		frame.StrokeColor = color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0x80}
		frame.StrokeWidth = 1
		return container.NewBorder(heading, nil, nil, nil,
			container.NewStack(frame, container.NewPadded(content)))
	}

	//docs:start overview
	w.SetContent(container.NewGridWithColumns(3,
		cell("NewVBox", vbox()),
		cell("NewHBox", hbox()),
		cell("NewBorder", border()),
		cell("NewGridWithColumns", grid()),
		cell("NewGridWrap", gridWrap()),
		cell("NewCenter", center()),
		cell("NewStack", stack()),
		cell("NewPadded", padded()),
		cell("FormLayout", form()),
	))
	//docs:end overview
	w.Resize(geo.NewSize(900, 650))
	bellui.ShowAndRun(w)
}
