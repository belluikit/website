// Play a video file or stream URL given on the command line:
//
//	go run . movie.mp4
//	go run . https://example.com/stream.m3u8
package main

import (
	"os"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/mmedia"
)

func main() {
	a := app.New()
	w := a.NewWindow("Player")
	w.Resize(geo.NewSize(800, 520))

	if len(os.Args) < 2 {
		w.SetContent(widget.NewLabel("Pass a media file or URL as the first argument."))
		bellui.ShowAndRun(w)
		return
	}

	//docs:start player
	status := widget.NewLabel("")
	mp, err := mmedia.NewMediaPlayerURL(os.Args[1],
		mmedia.WithAutoPlay(),
		mmedia.WithScrubPreview(), // show the frame under the pointer on the seek bar
		mmedia.WithOnError(func(err error) {
			// Callbacks run on the player's own goroutine.
			bellui.Do(func() { status.SetText(err.Error()) })
		}),
	)
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer mp.Close()
	//docs:end player

	w.SetContent(mp)
	bellui.ShowAndRun(w)
}
