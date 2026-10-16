// Screenshot program for the Media page: a MediaPlayer playing the sample
// video in this directory (a test pattern with a silent audio track).
//
//	cd examples/extensions/shots/player && go run .
//	go run . path/to/other.mp4
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

	path := "sample.mp4"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	mp, err := mmedia.NewMediaPlayerURL(path, mmedia.WithAutoPlay(), mmedia.WithLoop())
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer mp.Close()

	w.SetContent(mp)
	bellui.ShowAndRun(w)
}
