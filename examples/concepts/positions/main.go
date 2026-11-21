// Where an object is drawn: opening a pop-up just below a button, and skipping
// work for a label that is not on screen.
package main

import (
	"fmt"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.positions")
	w := a.NewWindow("Positions")

	//docs:start anchor
	var menu *widget.Button
	menu = widget.NewButton("Open below", func() {
		// The point just below the button's bottom-left corner, in window
		// coordinates. It follows any rotation or scale on the button or
		// its parents, so the pop-up opens where the button is drawn.
		below := bellui.AbsolutePositionForPoint(menu, geo.NewPos(0, menu.Size().Height))
		widget.ShowPopUpAtPosition(widget.NewLabel("A pop-up"), w.Canvas(), below)
	})
	//docs:end anchor

	//docs:start onscreen
	counter := widget.NewLabel("0")
	ticks := 0
	go func() {
		for range time.Tick(100 * time.Millisecond) {
			bellui.Do(func() {
				// Skip the work while the label's tab is not selected.
				if !bellui.IsOnScreen(counter) {
					return
				}
				ticks++
				counter.SetText(fmt.Sprint(ticks))
			})
		}
	}()

	tabs := container.NewAppTabs(
		container.NewTabItem("Counter", counter),
		container.NewTabItem("Other", widget.NewLabel("The counter pauses while this tab is shown.")),
	)
	//docs:end onscreen

	w.SetContent(container.NewBorder(container.NewHBox(menu), nil, nil, nil, tabs))
	w.Resize(geo.NewSize(400, 300))
	bellui.ShowAndRun(w)
}
