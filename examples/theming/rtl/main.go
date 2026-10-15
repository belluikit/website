// Right-to-left layout: forcing the direction for the whole app, and keeping
// one part of the window left-to-right.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Direction")

	row := container.NewHBox(
		widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), func() {}),
		widget.NewLabel("Leading"),
		widget.NewLabel("Trailing"),
	)

	//docs:start global
	direction := widget.NewRadioGroup([]string{"Auto", "Left to right", "Right to left"}, func(s string) {
		switch s {
		case "Left to right":
			bellui.SetLayoutDirection(bellui.DirectionLTR)
		case "Right to left":
			bellui.SetLayoutDirection(bellui.DirectionRTL)
		default:
			bellui.SetLayoutDirection(bellui.DirectionAuto) // follow the locale
		}
	})
	direction.SetSelected("Auto")
	//docs:end global

	//docs:start scope
	// This row always flows left to right, whatever the app direction is.
	code := container.NewDirectionScope(
		container.NewHBox(widget.NewLabel("x := 1"), widget.NewLabel("// comment")),
		bellui.DirectionLTR,
	)
	//docs:end scope

	w.SetContent(container.NewVBox(direction, row, widget.NewEntry(), code))
	bellui.ShowAndRun(w)
}
