// A wheel picker: spin through the hours of the day.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Wheel picker")
	status := widget.NewLabel("")

	//docs:start wheel
	var hours []core.CanvasObject
	for h := 0; h < 24; h++ {
		hours = append(hours, widget.NewLabel(fmt.Sprintf("%02d", h)))
	}
	wheel := widget.NewWheelPicker(hours, func(i int) {
		status.SetText(fmt.Sprintf("Hour: %02d", i))
	})
	wheel.Wrap = true // after 23 comes 00
	wheel.SetSelectedSilently(9)
	wheel.SetAccessibilityLabel("Hour")
	//docs:end wheel

	w.SetContent(container.NewVBox(container.NewCenter(wheel), status))
	bellui.ShowAndRun(w)
}
