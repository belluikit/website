// Binding to variables you own, to struct fields and to preferences.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/data/binding"
	"github.com/belluikit/bellui/widget"
)

//docs:start struct
type contact struct {
	Name  string
	Email string
}

//docs:end struct

func main() {
	a := app.NewWithID("com.example.binding")
	w := a.NewWindow("External data")

	//docs:start variable
	count := 0
	bound := binding.BindInt(&count) // reads and writes count

	label := widget.NewLabelWithData(binding.IntToString(bound))
	viaBinding := widget.NewButton("Add via binding", func() {
		v, _ := bound.Get()
		_ = bound.Set(v + 1) // updates count and the label
	})
	direct := widget.NewButton("Add directly", func() {
		count++
		_ = bound.Reload() // tell the binding that count changed
	})
	//docs:end variable

	//docs:start fields
	person := &contact{Name: "Ada", Email: "ada@example.com"}
	fields := binding.BindStruct(person)

	nameItem, _ := fields.GetItem("Name")
	emailItem, _ := fields.GetItem("Email")
	form := widget.NewForm(
		widget.NewFormItem("Name", widget.NewEntryWithData(nameItem.(binding.String))),
		widget.NewFormItem("Email", widget.NewEntryWithData(emailItem.(binding.String))),
	)
	//docs:end fields

	//docs:start prefs
	nickname := binding.BindPreferenceString("nickname", a.Preferences())
	nickEntry := widget.NewEntryWithData(nickname) // saved as you type
	//docs:end prefs

	w.SetContent(container.NewVBox(label, viaBinding, direct, form, nickEntry))
	bellui.ShowAndRun(w)
}
