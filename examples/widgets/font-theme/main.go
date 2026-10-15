// FontPicker and ThemeSelector for a settings page.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.settings")
	w := a.NewWindow("Settings")
	preview := widget.NewLabel("The quick brown fox")

	//docs:start font
	font := widget.NewFontPickerWithSystemFonts(
		widget.FontPickerConfig{Size: 14},
		func(cfg widget.FontPickerConfig) {
			preview.TextStyle = cfg.TextStyle()
			preview.SetText(fmt.Sprintf("%s, %.0f pt", cfg.Family, cfg.Size))
		})
	//docs:end font

	//docs:start theme
	appearance := widget.NewThemeSelector(func(sel widget.ThemeSelection) {
		sel.Apply()
		widget.SaveThemeSelection(sel, a.Preferences())
	})
	// Restore the choice saved last time.
	if sel, ok := widget.LoadThemeSelection(a.Preferences()); ok {
		sel.Apply()
		appearance.Selection = sel
	}
	//docs:end theme

	w.SetContent(container.NewVBox(appearance, widget.NewSeparator(), font, preview))
	bellui.ShowAndRun(w)
}
