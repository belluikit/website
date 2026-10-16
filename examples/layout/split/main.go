// A three-pane editor shape: a sidebar split from an editor on top, with two
// more panes stacked below it in a vertical Paned.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.split")
	w := a.NewWindow("Split containers")

	//docs:start split
	sidebar := widget.NewLabel("Sidebar")
	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Editor")

	// Two children with a draggable divider between them.
	split := container.NewHSplit(sidebar, editor)
	split.SetOffset(0.25) // the sidebar gets a quarter of the width
	//docs:end split

	//docs:start paned
	// Any number of children. Weights set their share of the space.
	panes := container.NewVPaned(
		split,
		widget.NewLabel("Output"),
		widget.NewLabel("Terminal"),
	)
	panes.SetWeight(0, 3) // the top pane gets 3 parts, the others 1 each
	//docs:end paned

	w.SetContent(panes)
	w.Resize(geo.NewSize(700, 500))
	bellui.ShowAndRun(w)
}
