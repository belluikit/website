// An app that goes back to its home screen when nobody has touched it for a
// while, built so that a test can check the idle rule.
package main

import (
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

//docs:start idle
const idleAfter = 2 * time.Minute // untouched this long, the app goes home

// isIdle reports whether the user has given no input for idleAfter.
// LastUserActivity is the zero time until the first input, so a fresh
// app counts as idle.
func isIdle(now time.Time) bool {
	return now.Sub(bellui.LastUserActivity()) >= idleAfter
}

//docs:end idle

func main() {
	a := app.New()
	w := a.NewWindow("Kiosk")

	home := widget.NewLabel("Welcome")
	details := widget.NewLabel("Details")
	stack := container.NewStack(home, details)
	details.Hide()
	open := widget.NewButton("Open details", func() {
		home.Hide()
		details.Show()
	})

	// Check once a second, and go home when idle.
	go func() {
		for range time.Tick(time.Second) {
			bellui.Do(func() {
				if isIdle(time.Now()) && details.Visible() {
					details.Hide()
					home.Show()
				}
			})
		}
	}()

	w.SetContent(container.NewVBox(stack, open))
	bellui.ShowAndRun(w)
}
