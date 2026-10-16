// Information, error and confirmation dialogs.
package main

import (
	"errors"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Dialogs")
	status := widget.NewLabel("Nothing chosen yet")

	//docs:start info
	info := widget.NewButton("Information", func() {
		dialog.ShowInformation("Saved", "Your changes were saved.", w)
	})
	failure := widget.NewButton("Error", func() {
		dialog.ShowError(errors.New("the disk is full"), w)
	})
	//docs:end info

	//docs:start confirm
	confirm := widget.NewButton("Delete…", func() {
		d := dialog.NewConfirm("Delete file?", "This cannot be undone.",
			func(ok bool) {
				if ok {
					status.SetText("Deleted")
				} else {
					status.SetText("Kept")
				}
			}, w)
		d.SetConfirmText("Delete")
		d.SetDismissText("Keep")
		d.SetConfirmImportance(widget.DangerImportance)
		d.Show()
	})
	//docs:end confirm

	//docs:start hosted
	hosted := widget.NewButton("In its own window", func() {
		d := dialog.NewInformation("Report ready", "The report has been generated.", w)
		d.SetWindowHosted(true) // a separate OS window on desktop
		d.SetOnClosed(func() { status.SetText("Report dialog closed") })
		d.Show()
	})
	//docs:end hosted

	w.SetContent(container.NewVBox(info, failure, confirm, hosted, status))
	bellui.ShowAndRun(w)
}
