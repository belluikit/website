// Subtitles built in code, from any source, instead of parsed from a file.
package main

import (
	"os"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/mmedia"
)

func main() {
	a := app.New()
	w := a.NewWindow("Cues")
	w.Resize(geo.NewSize(800, 520))

	if len(os.Args) < 2 {
		w.SetContent(widget.NewLabel("Pass a video file as the first argument."))
		bellui.ShowAndRun(w)
		return
	}

	mp, err := mmedia.NewMediaPlayerURL(os.Args[1], mmedia.WithAutoPlay())
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer mp.Close()

	//docs:start cues
	track := &mmedia.SubtitleTrack{
		Language: "en",
		Cues: []mmedia.SubtitleCue{
			{Start: 1 * time.Second, End: 4 * time.Second, Text: "Hello."},
			{Start: 5 * time.Second, End: 8 * time.Second, Text: "These cues were built in code."},
		},
	}
	subs := mmedia.NewSubtitleOverlay()
	subs.LoadTrack(track)
	w.SetContent(mmedia.MediaPlayerWithSubtitles(mp, subs))
	//docs:end cues

	bellui.ShowAndRun(w)
}
