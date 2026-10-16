// A clock updated from a background goroutine, and a button that starts slow
// work without freezing the window.
package main

import (
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.thread")
	w := a.NewWindow("UI thread")

	//docs:start clock
	clock := canvas.NewText("", theme.Color(theme.ColorNameForeground))
	clock.TextSize = 28

	go func() {
		for t := range time.Tick(time.Second) {
			// clock.Text is a plain field. Changing it and refreshing must
			// happen on the main goroutine, so hand the work to bellui.Do.
			bellui.Do(func() {
				clock.Text = t.Format("15:04:05")
				clock.Refresh()
			})
		}
	}()
	//docs:end clock

	//docs:start work
	status := widget.NewLabel("Idle")
	var start *widget.Button
	start = widget.NewButton("Start slow work", func() {
		// This callback runs on the main goroutine. Blocking here would
		// freeze the window, so the work moves to its own goroutine.
		start.Disable()
		status.SetText("Working...")
		go func() {
			time.Sleep(3 * time.Second) // stands in for real work

			bellui.Do(func() {
				status.SetText("Done")
				start.Enable()
			})
		}()
	})
	//docs:end work

	w.SetContent(container.NewVBox(container.NewCenter(clock), status, start))
	bellui.ShowAndRun(w)
}
