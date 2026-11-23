// Chips: assist, filter and input chips, and a group of chips to pick from.
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

	//docs:start group
	sizes := widget.NewChipGroup([]string{"Small", "Medium", "Large"}, false, func(picked []string) {
		status.SetText(fmt.Sprint("Size: ", picked))
	})
	sizes.Required = true
	sizes.Selected = []string{"Medium"}

	toppings := widget.NewChipGroup([]string{"Cheese", "Olives", "Peppers", "Onions"}, true, func(picked []string) {
		status.SetText(fmt.Sprint("Toppings: ", picked))
	})
	//docs:end group

	//docs:start exclude
	genres := widget.NewChipGroup([]string{"Rock", "Jazz", "Pop", "Folk"}, true, nil)
	genres.Excludable = true
	genres.OnStateChanged = func(option string, state widget.ChipState) {
		switch state {
		case widget.ChipStateSelected:
			status.SetText("Show " + option)
		case widget.ChipStateExcluded:
			status.SetText("Hide " + option)
		case widget.ChipStateOff:
			status.SetText(option + " no longer filtered")
		}
	}
	//docs:end exclude

	w.SetContent(container.NewVBox(
		container.NewHBox(directions, running),
		tags,
		sizes,
		toppings,
		genres,
		status,
	))
	bellui.ShowAndRun(w)
}
