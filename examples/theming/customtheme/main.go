// A custom theme that wraps the default theme and overrides a few colours,
// sizes, one icon and the fonts, plus a ThemeOverride for one part of the UI.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

//docs:start theme
type brandTheme struct {
	// The theme being customised. Every method brandTheme does not
	// define comes from it.
	bellui.Theme
}

var (
	_ bellui.Theme           = brandTheme{}
	_ theme.RenderStyleTheme = brandTheme{}
)

func newBrandTheme() brandTheme {
	return brandTheme{Theme: theme.DefaultTheme()}
}

func (t brandTheme) Color(n bellui.ThemeColorName, v bellui.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNamePrimary:
		if v == theme.VariantDark {
			return color.NRGBA{R: 0xff, G: 0xb3, B: 0x80, A: 0xff}
		}
		return color.NRGBA{R: 0xb3, G: 0x4a, B: 0x00, A: 0xff}
	case theme.ColorNameForegroundOnPrimary:
		if v == theme.VariantDark {
			return color.NRGBA{R: 0x3a, G: 0x18, B: 0x00, A: 0xff}
		}
		return color.White
	}
	return t.Theme.Color(n, v)
}

func (t brandTheme) Size(n bellui.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNameText:
		return 15
	case theme.SizeNameButtonRadius, theme.SizeNameInputRadius:
		return 0 // square corners
	}
	return t.Theme.Size(n)
}

func (t brandTheme) Icon(n bellui.ThemeIconName) bellui.Resource {
	if n == theme.IconNameConfirm {
		return t.Theme.Icon(theme.IconNameMailSend)
	}
	return t.Theme.Icon(n)
}

func (t brandTheme) Font(s bellui.TextStyle) bellui.Resource {
	// Return your own bundled font resources here; this keeps the defaults.
	return t.Theme.Font(s)
}

// RenderStyle keeps the default theme's modern widget look. Without it a
// wrapping theme falls back to the classic look.
func (t brandTheme) RenderStyle() theme.RenderStyle {
	if rs, ok := t.Theme.(theme.RenderStyleTheme); ok {
		return rs.RenderStyle()
	}
	return theme.RenderStyleModern
}

//docs:end theme

func main() {
	a := app.New()
	//docs:start apply
	a.Settings().SetTheme(newBrandTheme())
	//docs:end apply
	w := a.NewWindow("Custom theme")

	save := widget.NewButtonWithIcon("Send", theme.ConfirmIcon(), func() {})
	save.Importance = widget.HighImportance

	//docs:start override
	// A compact tool row: only this part of the window uses the dense theme.
	tools := container.NewThemeOverride(container.NewHBox(
		widget.NewButton("Cut", func() {}),
		widget.NewButton("Copy", func() {}),
		widget.NewButton("Paste", func() {}),
	), theme.DenseTheme())
	//docs:end override

	w.SetContent(container.NewVBox(
		widget.NewEntry(),
		save,
		tools,
	))
	bellui.ShowAndRun(w)
}
