// Chips: assist, filter and input chips.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Chips")
	status := widget.NewLabel("")

	//docs:start chips
	directions := widget.NewChip("Directions", func() {
		status.SetText("Directions")
	})
	directions.Icon = theme.NavigateNextIcon()

	running := widget.NewFilterChip("Running", func(on bool) {
		status.SetText(fmt.Sprint("Running: ", on))
	})
	running.Selected = true

	tags := container.NewHBox()
	tag := &widget.Chip{Text: "golang", Style: widget.ChipStyleInput}
	tag.ExtendBaseWidget(tag)
	tag.OnDeleted = func() {
		tags.Remove(tag)
	}
	tags.Add(tag)
	//docs:end chips

	w.SetContent(container.NewVBox(
		container.NewHBox(directions, running),
		tags,
		status,
	))
	bellui.ShowAndRun(w)
}
