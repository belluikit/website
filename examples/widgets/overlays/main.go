// Badges, snackbars and toasts.
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
	w := a.NewWindow("Messages")
	w.Resize(geo.NewSize(420, 320))
	items := []string{"a.txt", "b.txt"}
	list := widget.NewLabel("a.txt, b.txt")

	//docs:start badge
	inbox := widget.NewBadgeWithCount(widget.NewIconButton(theme.MailComposeIcon(), nil), 3)
	updates := widget.NewBadge(widget.NewIconButton(theme.DownloadIcon(), nil))
	//docs:end badge

	//docs:start snackbar
	archive := widget.NewButton("Archive", func() {
		removed := items
		items = nil
		list.SetText("(empty)")
		widget.ShowSnackbar("2 files archived", "Undo", func() {
			items = removed
			list.SetText("a.txt, b.txt")
		}, w.Canvas())
	})
	//docs:end snackbar

	//docs:start toast
	save := widget.NewButton("Save", func() {
		widget.ShowToastWithImportance("Saved", widget.SuccessImportance, w.Canvas())
	})
	//docs:end toast

	w.SetContent(container.NewVBox(container.NewHBox(inbox, updates), list, archive, save))
	bellui.ShowAndRun(w)
}
