// Context menus, pop-up menus and tooltips.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Menus")
	status := widget.NewLabel("Right-click the file name")

	//docs:start context
	file := widget.NewContextMenu(widget.NewLabel("report.pdf"),
		bellui.NewMenuItem("Open", func() { status.SetText("Open") }),
		bellui.NewMenuItem("Rename", func() { status.SetText("Rename") }),
		bellui.NewMenuItemSeparator(),
		bellui.NewMenuItem("Delete", func() { status.SetText("Delete") }),
	)
	//docs:end context

	//docs:start popup
	sortMenu := bellui.NewMenu("",
		bellui.NewMenuItem("By name", func() { status.SetText("Sort by name") }),
		bellui.NewMenuItem("By date", func() { status.SetText("Sort by date") }),
	)
	var sort *widget.Button
	sort = widget.NewButtonWithIcon("Sort", theme.MenuDropDownIcon(), func() {
		widget.ShowPopUpMenuAtRelativePosition(sortMenu, w.Canvas(),
			geo.NewPos(0, sort.Size().Height), sort)
	})
	//docs:end popup

	//docs:start tooltip
	refresh := widget.NewIconButton(theme.ViewRefreshIcon(), func() { status.SetText("Refreshed") })
	refresh.SetToolTipText("Reload the list")

	info := widget.NewTooltip(widget.NewIcon(theme.InfoIcon()), "Files are kept for 30 days")
	//docs:end tooltip

	w.SetContent(container.NewVBox(file, container.NewHBox(sort, refresh, info), status))
	bellui.ShowAndRun(w)
}
