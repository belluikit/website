// Data binding: widgets that share bound values, conversions between types,
// and updates from a background goroutine.
package main

import (
	"fmt"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/data/binding"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Data binding")

	//docs:start string
	name := binding.NewString()
	_ = name.Set("World")

	entry := widget.NewEntryWithData(name) // typing updates name
	greeting := widget.NewLabelWithData(
		binding.NewSprintf("Hello, %s!", name), // follows name
	)
	//docs:end string

	//docs:start float
	volume := binding.NewFloat()
	_ = volume.Set(0.3)

	slider := widget.NewSliderWithData(0, 1, volume)
	bar := widget.NewProgressBarWithData(volume)
	percent := widget.NewLabelWithData(
		binding.FloatToStringWithFormat(volume, "Volume: %.2f"),
	)
	//docs:end float

	//docs:start bool
	muted := binding.NewBool()
	mute := widget.NewCheckWithData("Mute", muted)
	loud := widget.NewCheckWithData("Sound on", binding.Not(muted))
	//docs:end bool

	//docs:start background
	clock := binding.NewString()
	go func() {
		for t := range time.Tick(time.Second) {
			_ = clock.Set(t.Format("15:04:05")) // safe from any goroutine
		}
	}()
	//docs:end background

	//docs:start listen
	changes := 0
	counter := widget.NewLabel("")
	onName := binding.NewDataListener(func() {
		changes++
		v, _ := name.Get()
		counter.SetText(fmt.Sprintf("%q, change %d", v, changes))
	})
	name.AddListener(onName)
	// Later, to stop listening: name.RemoveListener(onName)
	//docs:end listen

	w.SetContent(container.NewVBox(
		entry, greeting,
		slider, bar, percent,
		mute, loud,
		widget.NewLabelWithData(clock),
		counter,
	))
	bellui.ShowAndRun(w)
}
