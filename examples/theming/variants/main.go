// Switching the light/dark variant, the primary colour and the dense theme
// of a running app.
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
	w := a.NewWindow("Theme variants")

	//docs:start variant
	variant := widget.NewSegmentedButton([]string{"Light", "Dark"}, func(i int) {
		if i == 0 {
			a.Settings().SetThemeVariant(theme.VariantLight)
		} else {
			a.Settings().SetThemeVariant(theme.VariantDark)
		}
	})
	if a.Settings().ThemeVariant() == theme.VariantDark {
		variant.Selected = 1
	}
	//docs:end variant

	//docs:start primary
	primary := widget.NewSelect(theme.PrimaryColorNames(), func(name string) {
		if s, ok := a.Settings().(bellui.PrimaryColorSettings); ok {
			s.SetPrimaryColor(name)
		}
	})
	current := a.Settings().PrimaryColor()
	if current == "" {
		current = theme.ColorBlue // the default when nothing is set
	}
	primary.SetSelected(current)
	//docs:end primary

	//docs:start dense
	dense := widget.NewCheck("Dense theme", func(on bool) {
		if on {
			a.Settings().SetTheme(theme.DenseTheme())
		} else {
			a.Settings().SetTheme(theme.DefaultTheme())
		}
	})
	//docs:end dense

	save := widget.NewButton("Save", func() {})
	save.Importance = widget.HighImportance

	w.SetContent(container.NewVBox(
		variant,
		primary,
		dense,
		widget.NewEntry(),
		container.NewHBox(widget.NewButton("Cancel", func() {}), save),
	))

	bellui.ShowAndRun(w)
}
