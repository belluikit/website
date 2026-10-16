// Numeric selection: Slider, RangeSlider and SpinBox.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Sliders")
	status := widget.NewLabel("")

	//docs:start slider
	volume := widget.NewSlider(0, 100)
	volume.Step = 1
	volume.ShowValueIndicator = true
	volume.OnChanged = func(v float64) {
		status.SetText(fmt.Sprintf("Volume %.0f", v))
	}
	volume.OnChangeEnded = func(v float64) {
		status.SetText(fmt.Sprintf("Volume set to %.0f", v))
	}
	volume.SetValue(40)

	rating := widget.NewSlider(1, 5)
	rating.Step = 1
	rating.ShowTickMarks = true
	//docs:end slider

	//docs:start range
	price := widget.NewRangeSlider(0, 500)
	price.Step = 10
	price.SetValues(100, 300)
	price.OnChangeEnded = func(low, high float64) {
		status.SetText(fmt.Sprintf("Price %.0f to %.0f", low, high))
	}
	//docs:end range

	//docs:start spin
	guests := widget.NewSpinBox(1, 12, 1)
	guests.SetValue(2)
	guests.OnChanged = func(v float64) {
		status.SetText(fmt.Sprintf("%.0f guests", v))
	}
	//docs:end spin

	w.SetContent(container.NewVBox(volume, rating, price, guests, status))
	bellui.ShowAndRun(w)
}
