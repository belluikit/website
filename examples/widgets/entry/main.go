// Text entries: single line, outlined, password, multi-line and read-only.
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
	w := a.NewWindow("Entry")
	status := widget.NewLabel("")

	//docs:start basic
	name := widget.NewEntry()
	name.SetPlaceHolder("Your name")
	name.OnChanged = func(s string) {
		status.SetText("Typing: " + s)
	}
	name.OnSubmitted = func(s string) {
		status.SetText("Submitted: " + s)
	}
	//docs:end basic

	//docs:start variants
	email := widget.NewEntry()
	email.Style = widget.EntryStyleOutlined
	email.SetPlaceHolder("Email")
	email.SetIcon(theme.MailComposeIcon())

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	notes := widget.NewMultiLineEntry()
	notes.Wrapping = bellui.TextWrapWord
	notes.SetMinRowsVisible(4)

	licence := widget.NewMultiLineEntry()
	licence.SetText("Read-only text can be selected and copied.")
	licence.ReadOnly = true
	//docs:end variants

	w.SetContent(container.NewVBox(name, email, password, notes, licence, status))
	bellui.ShowAndRun(w)
}
