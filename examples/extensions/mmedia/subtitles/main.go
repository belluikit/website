// A video with SRT subtitles drawn over it.
package main

import (
	"os"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/mmedia"
)

const srt = `1
00:00:01,000 --> 00:00:04,000
Hello.

2
00:00:05,000 --> 00:00:08,000
These lines follow the playback time.
`

func main() {
	a := app.New()
	w := a.NewWindow("Subtitles")
	w.Resize(geo.NewSize(800, 520))

	if len(os.Args) < 2 {
		w.SetContent(widget.NewLabel("Pass a video file as the first argument."))
		bellui.ShowAndRun(w)
		return
	}

	//docs:start subtitles
	mp, err := mmedia.NewMediaPlayerURL(os.Args[1], mmedia.WithAutoPlay())
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer mp.Close()

	subs := mmedia.NewSubtitleOverlay()
	if err := subs.LoadSRT(strings.NewReader(srt)); err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	w.SetContent(mmedia.MediaPlayerWithSubtitles(mp, subs))
	//docs:end subtitles

	bellui.ShowAndRun(w)
}
