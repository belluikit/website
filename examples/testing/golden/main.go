// A status card whose rendering is checked against golden files.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/widget"
)

//docs:start ui
func statusCard(title, detail string) core.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle(title, bellui.TextAlignLeading, bellui.TextStyle{Bold: true}),
		widget.NewLabel(detail),
	)
}

//docs:end ui

func main() {
	a := app.New()
	w := a.NewWindow("Golden files")
	w.SetContent(statusCard("Backup", "Last run: today"))
	bellui.ShowAndRun(w)
}
