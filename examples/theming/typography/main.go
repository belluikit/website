// Text styles: theme text sizes on a Label, bold/italic/monospace, and the
// Material Design 3 typography presets on canvas.Text.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Typography")

	//docs:start label
	title := widget.NewLabel("Monthly report")
	title.SizeName = theme.SizeNameHeadlineSmall // 24 in the default theme
	title.TextStyle = bellui.TextStyle{Bold: true}

	note := widget.NewLabel("Figures are provisional.")
	note.TextStyle = bellui.TextStyle{Italic: true}

	code := widget.NewLabel("go run .")
	code.TextStyle = bellui.TextStyle{Monospace: true}
	//docs:end label

	//docs:start preset
	display := canvas.NewText("42 °C", theme.Color(theme.ColorNameForeground))
	display.TextSize = theme.Size(theme.SizeNameDisplayMedium) // 45
	display.TextStyle = theme.TypographyDisplayMedium

	medium := canvas.NewText("Medium weight", theme.Color(theme.ColorNameForeground))
	medium.TextSize = theme.Size(theme.SizeNameTitleMedium)
	medium.TextStyle = bellui.TextStyle{Weight: bellui.Ptr(bellui.WeightMedium)}
	//docs:end preset

	w.SetContent(container.NewVBox(title, note, code, display, medium))
	bellui.ShowAndRun(w)
}
