// Capture program for the pickers page: DateEntry and DateTimeEntry in a form.
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
	w := a.NewWindow("Date entries")

	birthday := widget.NewDateEntry()
	birthday.SetDate(time.Date(1990, time.July, 14, 0, 0, 0, 0, time.Local))
	meeting := widget.NewDateTimeEntry()
	meeting.SetDateTime(time.Date(2026, time.March, 9, 9, 41, 0, 0, time.Local))
	empty := widget.NewDateEntry()
	empty.SetPlaceHolder("Not set")

	w.SetContent(container.NewPadded(widget.NewForm(
		widget.NewFormItem("Birthday", birthday),
		widget.NewFormItem("Meeting", meeting),
		widget.NewFormItem("Deadline", empty),
	)))
	bellui.ShowAndRun(w)
}
