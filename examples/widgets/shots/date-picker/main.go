// Capture program for the pickers page: a DatePicker shown in place.
package main

import (
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Date picker")
	day := time.Date(2026, time.March, 9, 0, 0, 0, 0, time.Local)
	w.SetContent(container.NewCenter(widget.NewDatePicker(day, nil)))
	bellui.ShowAndRun(w)
}
