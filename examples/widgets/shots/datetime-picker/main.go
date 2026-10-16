// Capture program for the pickers page: a DateTimePicker, date and time on two tabs.
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
	w := a.NewWindow("Date and time picker")
	when := time.Date(2026, time.March, 9, 9, 41, 0, 0, time.Local)
	w.SetContent(container.NewCenter(widget.NewDateTimePicker(when, false, nil)))
	bellui.ShowAndRun(w)
}
