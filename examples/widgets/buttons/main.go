// Buttons: the MD3 styles, importance, icons, button groups, segmented and
// split buttons.
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
	w := a.NewWindow("Buttons")
	status := widget.NewLabel("Tap something")
	say := func(s string) func() { return func() { status.SetText(s) } }

	//docs:start basic
	save := widget.NewButton("Save", func() {
		status.SetText("Saved")
	})
	open := widget.NewButtonWithIcon("Open", theme.FolderOpenIcon(), say("Open"))
	//docs:end basic

	//docs:start styles
	filled := widget.NewButton("Filled", say("Filled"))
	filled.Style = widget.ButtonStyleFilled
	tonal := widget.NewButton("Tonal", say("Tonal"))
	tonal.Style = widget.ButtonStyleFilledTonal
	elevated := widget.NewButton("Elevated", say("Elevated"))
	elevated.Style = widget.ButtonStyleElevated
	outlined := widget.NewButton("Outlined", say("Outlined"))
	outlined.Style = widget.ButtonStyleOutlined
	text := widget.NewButton("Text", say("Text"))
	text.Style = widget.ButtonStyleText
	//docs:end styles

	//docs:start importance
	del := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), say("Deleted"))
	del.Importance = widget.DangerImportance

	next := widget.NewButtonWithIcon("Next", theme.NavigateNextIcon(), say("Next"))
	next.Importance = widget.HighImportance
	next.IconPlacement = widget.ButtonIconTrailingText
	//docs:end importance

	//docs:start icon
	settings := widget.NewIconButton(theme.SettingsIcon(), say("Settings"))
	settings.SetAccessibilityLabel("Settings")
	settings.SetToolTipText("Settings")

	add := widget.NewIconButton(theme.ContentAddIcon(), say("Add"))
	add.Style = widget.IconButtonStyleFilled
	//docs:end icon

	//docs:start group
	group := widget.NewButtonGroup(
		widget.NewButton("Cut", say("Cut")),
		widget.NewButton("Copy", say("Copy")),
		widget.NewButton("Paste", say("Paste")),
	)

	period := widget.NewSegmentedButton([]string{"Day", "Week", "Month"}, func(i int) {
		status.SetText(fmt.Sprintf("Segment %d", i))
	})
	period.Selected = 1

	split := widget.NewSplitButton("Export", say("Export"), []*bellui.MenuItem{
		bellui.NewMenuItem("Export as PDF", say("PDF")),
		bellui.NewMenuItem("Export as CSV", say("CSV")),
	})
	//docs:end group

	w.SetContent(container.NewVBox(
		container.NewHBox(save, open),
		container.NewHBox(filled, tonal, elevated, outlined, text),
		container.NewHBox(del, next),
		container.NewHBox(settings, add),
		group,
		period,
		container.NewHBox(split),
		status,
	))
	bellui.ShowAndRun(w)
}
