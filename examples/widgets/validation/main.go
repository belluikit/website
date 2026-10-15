// Validating an Entry with a regular expression and a custom rule.
package main

import (
	"errors"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/data/validation"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Validation")

	//docs:start validate
	zip := widget.NewEntry()
	zip.SetPlaceHolder("Postcode (5 digits)")
	zip.Validator = validation.NewRegexp(`^\d{5}$`, "must be five digits")

	user := widget.NewEntry()
	user.SetPlaceHolder("Username")
	user.Validator = func(s string) error {
		if strings.ContainsAny(s, " \t") {
			return errors.New("no spaces")
		}
		return nil
	}
	//docs:end validate

	status := widget.NewLabel("")
	//docs:start check
	check := widget.NewButton("Check", func() {
		if err := user.Validate(); err != nil {
			status.SetText("Username " + err.Error())
			return
		}
		status.SetText("Looks good")
	})
	//docs:end check

	w.SetContent(container.NewVBox(zip, user, check, status))
	bellui.ShowAndRun(w)
}
