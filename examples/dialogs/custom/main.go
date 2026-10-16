// Custom dialogs and form dialogs.
package main

import (
	"errors"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Custom dialogs")
	status := widget.NewLabel("")

	//docs:start custom
	custom := widget.NewButton("Custom content", func() {
		content := container.NewVBox(
			widget.NewLabel("Any canvas object can be the body."),
			widget.NewCheck("Show this again", nil),
		)
		dialog.ShowCustom("Tip of the day", "Close", content, w)
	})
	//docs:end custom

	//docs:start customconfirm
	rename := widget.NewButton("Rename…", func() {
		name := widget.NewEntry()
		name.SetPlaceHolder("New name")
		dialog.ShowCustomConfirm("Rename", "Rename", "Cancel", name,
			func(ok bool) {
				if ok {
					status.SetText("Renamed to " + name.Text)
				}
			}, w)
	})
	//docs:end customconfirm

	//docs:start buttons
	choose := widget.NewButton("Own buttons", func() {
		var d *dialog.CustomDialog
		pick := func(size string) func() {
			return func() {
				status.SetText("Size: " + size)
				d.Hide()
			}
		}
		d = dialog.NewCustomWithoutButtons("Pick a size",
			widget.NewLabel("Which size do you want?"), w)
		d.SetButtons([]core.CanvasObject{
			widget.NewButton("Small", pick("small")),
			widget.NewButton("Medium", pick("medium")),
			widget.NewButton("Large", pick("large")),
		})
		d.Show()
	})
	//docs:end buttons

	//docs:start form
	signUp := widget.NewButton("Sign up…", func() {
		email := widget.NewEntry()
		email.Validator = func(s string) error {
			if !strings.Contains(s, "@") {
				return errors.New("enter an email address")
			}
			return nil
		}
		password := widget.NewPasswordEntry()

		items := []*widget.FormItem{
			widget.NewFormItem("Email", email),
			widget.NewFormItem("Password", password),
		}
		dialog.ShowForm("Create account", "Sign up", "Cancel", items,
			func(ok bool) {
				if ok {
					status.SetText("Signed up as " + email.Text)
				}
			}, w)
	})
	//docs:end form

	w.SetContent(container.NewVBox(custom, rename, choose, signUp, status))
	bellui.ShowAndRun(w)
}
