// Device information, orientation and app lifecycle on mobile.
package main

import (
	"strconv"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/driver/mobile"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.mobile")
	w := a.NewWindow("Device")

	info := widget.NewLabel("")
	events := widget.NewLabel("")

	//docs:start device
	describe := func() {
		d := bellui.CurrentDevice()
		shape := "portrait"
		if bellui.IsHorizontal(d.Orientation()) {
			shape = "landscape"
		}
		info.SetText("Mobile: " + strconv.FormatBool(d.IsMobile()) +
			"\nBrowser: " + strconv.FormatBool(d.IsBrowser()) +
			"\nHardware keyboard: " + strconv.FormatBool(d.HasKeyboard()) +
			"\nOrientation: " + shape)
	}
	//docs:end device

	//docs:start lifecycle
	life := a.Lifecycle()
	life.SetOnStarted(describe)
	life.SetOnEnteredForeground(func() {
		describe() // the device may have been rotated while we were away
		events.SetText("Back in the foreground")
	})
	life.SetOnExitedForeground(func() {
		events.SetText("Went to the background")
	})
	//docs:end lifecycle

	//docs:start keyboard
	hide := widget.NewButton("Hide keyboard", func() {
		if d, ok := bellui.CurrentDevice().(mobile.Device); ok {
			d.HideVirtualKeyboard()
		}
	})
	//docs:end keyboard

	w.SetContent(container.NewVBox(info, events, widget.NewEntry(), hide))
	bellui.ShowAndRun(w)
}
