// An idle timer: after ten seconds without input the window shows a notice,
// and the next input brings the content back.
package main

import (
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.idle")
	w := a.NewWindow("Idle timer")

	content := container.NewVBox(widget.NewLabel("Type, click or scroll."), widget.NewEntry())
	idle := widget.NewLabel("Nobody here? Touch anything to come back.")
	idle.Hide()

	//docs:start idle
	const idleAfter = 10 * time.Second
	started := time.Now()
	go func() {
		for range time.Tick(time.Second) {
			bellui.Do(func() {
				// LastUserActivity is the zero time until the first input,
				// so count from the start of the app until then.
				last := bellui.LastUserActivity()
				if last.Before(started) {
					last = started
				}
				away := time.Since(last) > idleAfter
				if away && content.Visible() {
					content.Hide()
					idle.Show()
				} else if !away && !content.Visible() {
					idle.Hide()
					content.Show()
				}
			})
		}
	}()
	//docs:end idle

	w.SetContent(container.NewStack(content, container.NewCenter(idle)))
	w.Resize(geo.NewSize(360, 200))
	bellui.ShowAndRun(w)
}
