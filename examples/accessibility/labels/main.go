// Naming widgets that have no visible text, and announcing status changes
// to a screen reader.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/lang"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Labels")

	status := widget.NewLabel("3 messages")

	//docs:start label
	del := widget.NewIconButton(theme.DeleteIcon(), func() {
		status.SetText("Messages deleted")
		//docs:start announce
		bellui.AnnounceToScreenReader(lang.L("Messages deleted"), false)
		//docs:end announce
	})
	del.SetAccessibilityLabel(lang.L("Delete messages"))
	//docs:end label

	//docs:start text
	// A button with text needs no extra label: its text is its name.
	archive := widget.NewButton(lang.L("Archive"), func() {})
	//docs:end text

	//docs:start form
	// A form's visible labels are not linked to its fields, and an entry is
	// named by its placeholder. Name entries in a form explicitly.
	email := widget.NewEntry()
	email.SetAccessibilityLabel(lang.L("Email"))
	form := widget.NewForm(widget.NewFormItem(lang.L("Email"), email))
	//docs:end form

	w.SetContent(container.NewVBox(status, container.NewHBox(del, archive), form))
	bellui.ShowAndRun(w)
}
