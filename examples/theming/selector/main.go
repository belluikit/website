// An appearance setting built from widget.ThemeSelector, applied live and
// remembered in the app's own preferences.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/widget"
)

func main() {
	//docs:start startup
	// Preferences need an app ID.
	a := app.NewWithID("com.example.themeselector")

	// Restore the saved appearance before the first window is shown.
	if saved, ok := widget.LoadThemeSelection(a.Preferences()); ok {
		saved.Apply()
	}
	//docs:end startup

	w := a.NewWindow("Appearance")

	//docs:start selector
	selector := widget.NewThemeSelector(func(sel widget.ThemeSelection) {
		sel.Apply()                                     // re-theme this app now
		widget.SaveThemeSelection(sel, a.Preferences()) // remember it
	})
	if saved, ok := widget.LoadThemeSelection(a.Preferences()); ok {
		selector.Selection = saved // show the saved choice without firing OnChanged
	}
	//docs:end selector

	w.SetContent(widget.NewCard("Appearance", "Theme and accent colour", selector))
	bellui.ShowAndRun(w)
}
