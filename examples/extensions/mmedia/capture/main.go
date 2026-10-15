// Show a camera, with a picker to switch between cameras.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/mmedia"
)

func main() {
	a := app.New()
	w := a.NewWindow("Camera")
	w.Resize(geo.NewSize(800, 560))

	//docs:start capture
	cam, err := mmedia.NewCapturePlayer("", mmedia.WithAutoPlay()) // "" is the default camera
	if err != nil {
		w.SetContent(widget.NewLabel(err.Error()))
		bellui.ShowAndRun(w)
		return
	}
	defer cam.Close()

	status := widget.NewLabel("")
	picker := mmedia.NewCaptureSelect(func(d mmedia.CaptureDevice) {
		// Opening a camera takes a while: do it off the UI thread.
		go func() {
			if err := cam.SetCaptureDevice(d.Name); err != nil {
				bellui.Do(func() { status.SetText(err.Error()) })
			}
		}()
	})
	//docs:end capture

	w.SetContent(container.NewBorder(container.NewHBox(picker, status), nil, nil, nil, cam))
	bellui.ShowAndRun(w)
}
