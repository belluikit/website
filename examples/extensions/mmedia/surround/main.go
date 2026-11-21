// Ask for 5.1 output before the first player opens the audio device.
package main

import (
	"os"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/mmedia"
)

func main() {
	a := app.New()
	w := a.NewWindow("Surround")

	if len(os.Args) < 2 {
		w.SetContent(widget.NewLabel("Pass an audio file as the first argument."))
		bellui.ShowAndRun(w)
		return
	}

	//docs:start surround
	// Before any player with audio exists. Without this call the device is stereo.
	if err := mmedia.SetAudioDeviceChannels(6); err != nil { // 6 for 5.1, 8 for 7.1
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	ap, err := mmedia.NewAudioPlayerURL(os.Args[1], mmedia.WithAutoPlay())
	//docs:end surround
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer ap.Close()

	w.SetContent(ap)
	bellui.ShowAndRun(w)
}
