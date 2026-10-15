// Selection controls: Check, CheckGroup, RadioGroup, Select and Switch.
package main

import (
	"fmt"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Selection")
	status := widget.NewLabel("")

	//docs:start check
	terms := widget.NewCheck("I accept the terms", func(on bool) {
		status.SetText(fmt.Sprint("Accepted: ", on))
	})

	toppings := widget.NewCheckGroup([]string{"Cheese", "Olives", "Basil"}, func(sel []string) {
		status.SetText("Toppings: " + strings.Join(sel, ", "))
	})
	toppings.Horizontal = true
	toppings.SetSelected([]string{"Cheese"})
	//docs:end check

	//docs:start radio
	size := widget.NewRadioGroup([]string{"Small", "Medium", "Large"}, func(s string) {
		status.SetText("Size: " + s)
	})
	size.Required = true
	size.SetSelected("Medium")
	//docs:end radio

	//docs:start select
	lang := widget.NewSelect([]string{"English", "Deutsch", "Français"}, func(s string) {
		status.SetText("Language: " + s)
	})
	lang.PlaceHolder = "Language"
	//docs:end select

	//docs:start switch
	wifi := widget.NewSwitch(func(on bool) {
		status.SetText(fmt.Sprint("Wi-Fi: ", on))
	})
	wifi.On = true
	//docs:end switch

	w.SetContent(container.NewVBox(
		terms, toppings, size, lang,
		container.NewHBox(widget.NewLabel("Wi-Fi"), wifi),
		status,
	))
	bellui.ShowAndRun(w)
}
