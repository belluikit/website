// Disabling widgets: one switch disables a whole set of tabs and what is
// inside them.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Disabled")

	//docs:start disable
	name := widget.NewEntry()
	name.SetText("Ada")
	save := widget.NewButton("Save", func() {})
	sizes := widget.NewChipGroup([]string{"S", "M", "L"}, false, nil)

	tabs := container.NewAppTabs(
		container.NewTabItem("Profile", container.NewVBox(name, save)),
		container.NewTabItem("Size", sizes),
	)

	// The app's own choice: Save stays disabled until something changes.
	save.Disable()
	name.OnChanged = func(string) { save.Enable() }

	lock := widget.NewCheck("Lock", func(on bool) {
		if on {
			tabs.Disable() // disables the entry, the button and the chips too
		} else {
			tabs.Enable() // enables only what Disable switched off
		}
	})
	//docs:end disable

	w.SetContent(container.NewBorder(lock, nil, nil, nil, tabs))
	bellui.ShowAndRun(w)
}
