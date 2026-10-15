// Progress and loading: ProgressBar, ProgressBarInfinite, Activity and
// LoadingIndicator.
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
	w := a.NewWindow("Progress")

	//docs:start progress
	download := widget.NewProgressBar()
	download.TextFormatter = func() string {
		return fmt.Sprintf("%.0f of 100 MB", download.Value*100)
	}
	go func() {
		for i := 0; i <= 100; i++ {
			time.Sleep(50 * time.Millisecond)
			download.SetValue(float64(i) / 100) // safe from any goroutine
		}
	}()

	waiting := widget.NewProgressBarInfinite()
	stop := widget.NewButton("Stop", waiting.Stop)
	//docs:end progress

	//docs:start activity
	spinner := widget.NewActivity()
	spinner.Start()

	avatar := widget.NewLoadingIndicator(widget.LoadingShapeCircle)
	line := widget.NewLoadingIndicator(widget.LoadingShapeText)
	//docs:end activity

	w.SetContent(container.NewVBox(
		download, waiting, stop,
		container.NewHBox(spinner, widget.NewLabel("Syncing…")),
		container.NewHBox(avatar, container.NewVBox(line)),
	))
	bellui.ShowAndRun(w)
}
