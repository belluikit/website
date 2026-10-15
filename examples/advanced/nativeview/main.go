// The NativeView contract, shown with a view that only logs what the host asks.
package main

import (
	"log"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

//docs:start view
// loggingView stands in for a real platform view (a map, a video surface).
// A real backend would create the OS view in Attach and move it in SetGeometry.
type loggingView struct{}

func (loggingView) Attach(win bellui.Window) error {
	log.Println("attach to", win.Title())
	return nil
}

func (loggingView) SetGeometry(rect, clip geo.Rectangle, scale float32) {
	log.Printf("place at %v, clip to %v, scale %.1f", rect, clip, scale)
}

func (loggingView) SetVisible(visible bool) {
	log.Println("visible:", visible)
}

func (loggingView) Destroy() {
	log.Println("destroy")
}

// Check at compile time that loggingView satisfies the contract.
var _ widget.NativeView = loggingView{}

//docs:end view

func main() {
	a := app.New()
	w := a.NewWindow("Native view")

	//docs:start host
	host := widget.NewNativeViewHost(loggingView{})
	host.SetMinSize(geo.NewSize(300, 200))
	//docs:end host

	toggle := widget.NewButton("Hide / show", func() {
		if host.Visible() {
			host.Hide()
		} else {
			host.Show()
		}
	})

	w.SetContent(container.NewBorder(toggle, nil, nil, nil, host))
	bellui.ShowAndRun(w)
}
