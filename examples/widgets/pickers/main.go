// Date and time pickers, inline and as modal dialogs.
package main

import (
	"fmt"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Pickers")
	status := widget.NewLabel("")

	//docs:start inline
	calendar := widget.NewDatePicker(time.Now(), func(d time.Time) {
		status.SetText("Date: " + d.Format("2006-01-02"))
	})
	//docs:end inline

	//docs:start modal
	pickTime := widget.NewButton("Pick a time", func() {
		now := time.Now()
		widget.ShowAutoTimePicker(now.Hour(), now.Minute(), func(h, m int) {
			status.SetText(fmt.Sprintf("Time: %02d:%02d", h, m))
		}, w.Canvas())
	})

	pickBoth := widget.NewButton("Pick date and time", func() {
		widget.ShowAutoDateTimePicker(time.Now(), func(t time.Time) {
			status.SetText("When: " + t.Format(time.RFC822))
		}, w.Canvas())
	})
	//docs:end modal

	w.SetContent(container.NewVBox(calendar, pickTime, pickBoth, status))
	bellui.ShowAndRun(w)
}
