// Capture program for the pickers page: a TimePicker clock dial in 12-hour mode.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Time picker")
	w.SetContent(container.NewCenter(widget.NewTimePicker(9, 41, false, nil)))
	bellui.ShowAndRun(w)
}
