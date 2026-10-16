// Form: labelled rows, hints, validation and submit/cancel buttons.
package main

import (
	"errors"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/data/validation"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Sign up")
	status := widget.NewLabel("")

	//docs:start form
	name := widget.NewEntry()
	name.Validator = func(s string) error {
		if s == "" {
			return errors.New("enter your name")
		}
		return nil
	}

	email := widget.NewEntry()
	email.Validator = validation.NewRegexp(`^[^@\s]+@[^@\s]+\.[a-z]+$`, "not a valid email")

	plan := widget.NewSelect([]string{"Free", "Pro"}, nil)
	plan.SetSelected("Free")

	newsletter := widget.NewCheck("Send me the newsletter", nil)

	form := widget.NewForm(
		widget.NewFormItem("Name", name),
		&widget.FormItem{Text: "Email", Widget: email, HintText: "We never share it"},
		widget.NewFormItem("Plan", plan),
	)
	form.Append("", newsletter)
	form.SubmitText = "Create account"
	form.OnSubmit = func() {
		status.SetText("Welcome, " + name.Text)
	}
	form.OnCancel = func() {
		w.Close()
	}
	//docs:end form

	w.SetContent(container.NewVBox(form, status))
	bellui.ShowAndRun(w)
}
