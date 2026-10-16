// A long list of labels in a vertical scroll, with buttons that jump to the
// ends and a label that follows the scroll position.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.scroll")
	w := a.NewWindow("Scrolling")

	//docs:start scroll
	rows := container.NewVBox()
	for i := 1; i <= 100; i++ {
		rows.Add(widget.NewLabel(fmt.Sprintf("Row %d", i)))
	}

	// Only vertical scrolling: the content is as wide as the scroll.
	scroll := container.NewVScroll(rows)
	//docs:end scroll

	//docs:start control
	where := widget.NewLabel("Offset 0")
	show := func(offset geo.Position) {
		where.SetText(fmt.Sprintf("Offset %.0f", offset.Y))
	}
	// Called when the user scrolls, and by ScrollToBottom.
	scroll.OnScrolled = show

	top := widget.NewButton("Top", func() {
		scroll.ScrollToTop() // does not call OnScrolled
		show(scroll.Offset)
	})
	bottom := widget.NewButton("Bottom", scroll.ScrollToBottom)
	//docs:end control

	w.SetContent(container.NewBorder(
		container.NewHBox(top, bottom, where), nil, nil, nil,
		scroll,
	))
	w.Resize(geo.NewSize(320, 400))
	bellui.ShowAndRun(w)
}
