// Developer tools: the F10/F11/F12 overlay and the widget-tree inspector.
package main

import (
	"os"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dev"
	"github.com/belluikit/bellui/dev/inspector"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Dev tools")
	w.SetContent(container.NewPadded(container.NewVBox(
		widget.NewLabel("Press F12, F11 or F10"),
		widget.NewEntry(),
		container.NewHBox(widget.NewButton("OK", func() {}), widget.NewButton("Cancel", func() {})),
	)))

	//docs:start enable
	// F12: outline overlay, F11: print the tree, F10: refresh tracing.
	dev.Enable(w)
	//docs:end enable

	//docs:start inspector
	// Opens a second window that inspects w. Run with BELLUI_INSPECT=1.
	if os.Getenv("BELLUI_INSPECT") == "1" {
		inspector.Show(w)
	}
	//docs:end inspector

	bellui.ShowAndRun(w)
}
