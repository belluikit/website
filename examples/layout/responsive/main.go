// A dashboard that reflows with the window width: one card per row on a
// narrow window, two on a medium one, four on a wide one.
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

func card(title string) core.CanvasObject {
	bg := canvas.NewRectangle(color.NRGBA{R: 0x3b, G: 0x6e, B: 0xc8, A: 0x50})
	bg.CornerRadius = 8
	bg.SetMinSize(geo.NewSize(0, 80))
	return container.NewStack(bg, container.NewCenter(widget.NewLabel(title)))
}

func main() {
	a := app.NewWithID("com.example.responsive")
	w := a.NewWindow("Responsive")

	//docs:start responsive
	var cards []core.CanvasObject
	for i := 1; i <= 4; i++ {
		// Width ratios per size class: compact, medium, expanded.
		// Classes after the last ratio given reuse it.
		cards = append(cards, layout.Responsive(card(fmt.Sprintf("Card %d", i)), 1, 0.5, 0.25))
	}
	// A plain object takes a full row in every size class.
	cards = append(cards, card("Full-width footer"))

	grid := layout.NewResponsiveLayout(cards...)
	//docs:end responsive

	//docs:start breakpoints
	// Use your own size-class boundaries instead of the defaults.
	grid.Layout.(*layout.ResponsiveLayout).Breakpoints = []float32{576, 768, 992, 1200}
	//docs:end breakpoints

	//docs:start adaptive
	// Three columns while the device is landscape, three rows in portrait.
	adaptive := container.NewAdaptiveGrid(3, card("A"), card("B"), card("C"))
	//docs:end adaptive

	//docs:start rowwrap
	// Each child keeps its own width; rows wrap when the width runs out.
	tags := container.New(layout.NewRowWrapLayout())
	for _, t := range []string{"go", "ui", "desktop", "mobile", "web", "webgpu", "layout"} {
		tags.Add(widget.NewButton(t, func() {}))
	}
	//docs:end rowwrap

	w.SetContent(container.NewVScroll(container.NewVBox(grid, adaptive, tags)))
	w.Resize(geo.NewSize(900, 400))
	bellui.ShowAndRun(w)
}
