// A theme defined as JSON: only the listed values change, everything else
// comes from the default theme.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

//docs:start json
const themeJSON = `{
  "Colors": {
    "primary": "#6a1b9a"
  },
  "Colors-dark": {
    "primary": "#ce93d8",
    "background": "#121016"
  },
  "Sizes": {
    "text": 15,
    "padding": 8
  }
}`

//docs:end json

func main() {
	a := app.New()

	//docs:start load
	th, err := theme.FromJSON(themeJSON)
	if err != nil {
		bellui.LogError("Theme JSON is invalid, using the default theme", err)
	}
	a.Settings().SetTheme(th) // th is the default theme when err != nil
	//docs:end load

	w := a.NewWindow("JSON theme")
	go1 := widget.NewButton("Go", func() {})
	go1.Importance = widget.HighImportance
	w.SetContent(container.NewVBox(widget.NewLabel("Purple accents"), widget.NewEntry(), go1))
	bellui.ShowAndRun(w)
}
