// Play an audio file through an io.Reader, with a progress label of our own.
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
	w := a.NewWindow("Audio")

	if len(os.Args) < 2 {
		w.SetContent(widget.NewLabel("Pass an audio file as the first argument."))
		bellui.ShowAndRun(w)
		return
	}

	//docs:start audio
	f, err := os.Open(os.Args[1]) // *os.File is seekable, so the seek bar works
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer f.Close()

	progress := widget.NewLabel("")
	ap, err := mmedia.NewAudioPlayer(f,
		mmedia.WithVolume(0.8),
		mmedia.WithOnTimeUpdate(func(current, total time.Duration) {
			bellui.Do(func() {
				progress.SetText(fmt.Sprintf("%s of %s", current.Round(time.Second), total.Round(time.Second)))
			})
		}),
	)
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer ap.Close()
	ap.Play()
	//docs:end audio

	w.SetContent(container.NewVBox(ap, progress))
	bellui.ShowAndRun(w)
}
