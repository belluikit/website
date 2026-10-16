// DateEntry and DateTimeEntry: typed, locale-formatted values with a picker button.
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
	status := widget.NewLabel("")

	//docs:start entries
	birthday := widget.NewDateEntry()
	birthday.SetPlaceHolder("Birthday")
	birthday.OnChanged = func(d time.Time) {
		status.SetText("Birthday " + d.Format("2 Jan 2006"))
	}

	meeting := widget.NewDateTimeEntry()
	meeting.SetDateTime(time.Now().Add(24 * time.Hour))
	meeting.OnChanged = func(t time.Time) {
		status.SetText("Meeting " + t.Format(time.Kitchen))
	}
	//docs:end entries

	w.SetContent(container.NewVBox(birthday, meeting, status))
	bellui.ShowAndRun(w)
}
