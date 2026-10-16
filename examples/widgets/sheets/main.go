// Bottom and side sheets, and a plain pop-up.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Sheets")
	w.Resize(geo.NewSize(500, 400))

	//docs:start bottom
	showShare := widget.NewButton("Share…", func() {
		sheet := widget.NewBottomSheet(nil)
		sheet.Modal = true
		sheet.Content = container.NewVBox(
			widget.NewLabel("Share to"),
			widget.NewButton("Copy link", func() { sheet.Hide() }),
			widget.NewButton("Cancel", func() { sheet.Hide() }),
		)
		sheet.Show(w.Canvas())
	})
	//docs:end bottom

	//docs:start side
	showDetails := widget.NewButton("Details", func() {
		widget.ShowSideSheet(container.NewVBox(
			widget.NewLabel("Size: 2.4 MB"),
			widget.NewLabel("Modified: today"),
		), w.Canvas())
	})
	//docs:end side

	//docs:start popup
	var pop *widget.PopUp
	pop = widget.NewModalPopUp(container.NewVBox(
		widget.NewLabel("Are you still there?"),
		widget.NewButton("Yes", func() { pop.Hide() }),
	), w.Canvas())
	showPop := widget.NewButton("Pop-up", func() { pop.Show() })
	//docs:end popup

	w.SetContent(container.NewVBox(showShare, showDetails, showPop))
	bellui.ShowAndRun(w)
}
