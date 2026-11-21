// Audio that keeps playing while the screen showing it is replaced.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/mmedia"
)

func main() {
	a := app.New()
	w := a.NewWindow("Playback")

	if len(os.Args) < 2 {
		w.SetContent(widget.NewLabel("Pass an audio file or URL as the first argument."))
		bellui.ShowAndRun(w)
		return
	}

	//docs:start playback
	seeks := widget.NewLabel("")
	pb, err := mmedia.NewAudioPlaybackURL(os.Args[1],
		mmedia.WithOnSeek(func(from, to time.Duration) {
			bellui.Do(func() {
				seeks.SetText(fmt.Sprintf("Seek from %s to %s", from.Round(time.Second), to.Round(time.Second)))
			})
		}),
	)
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer pb.Close() // the app owns the playback, so the app closes it
	pb.Play()

	// Each screen gets its own view of the same playback. Closing a view
	// leaves the song playing.
	var view *mmedia.AudioPlayer
	var showScreen func(n int)
	showScreen = func(n int) {
		if view != nil {
			view.Close()
		}
		view = mmedia.NewAudioPlayerFor(pb)
		next := widget.NewButton("Next screen", func() { showScreen(n + 1) })
		w.SetContent(container.NewVBox(
			widget.NewLabel(fmt.Sprintf("Screen %d", n)),
			view, seeks, next,
		))
	}
	showScreen(1)
	//docs:end playback

	bellui.ShowAndRun(w)
}
