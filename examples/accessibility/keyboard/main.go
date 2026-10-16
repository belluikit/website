// Keyboard focus: where it starts, and the order Tab follows.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Keyboard")

	//docs:start order
	// The focus order is the order in the tree: name, email, notes, then the
	// button. The disabled check is skipped. Tab stops at notes: a multi-line
	// entry keeps Tab and Shift+Tab, so the keyboard cannot move focus past it.
	name := widget.NewEntry()
	email := widget.NewEntry()
	notes := widget.NewMultiLineEntry() // takes Tab itself, to insert a tab
	later := widget.NewCheck("Remind me later", nil)
	later.Disable()
	send := widget.NewButton("Send", func() {})

	w.SetContent(container.NewVBox(name, email, notes, later, send))
	//docs:end order

	//docs:start initial
	w.Canvas().Focus(name) // start typing straight away
	//docs:end initial

	bellui.ShowAndRun(w)
}
