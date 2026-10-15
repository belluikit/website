// A theme whose whole Material Design 3 colour scheme is generated from one
// seed colour that is not among the named primary colours.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

//docs:start seed
type seedTheme struct {
	bellui.Theme                  // sizes, fonts and icons
	light, dark  *theme.MD3Scheme // colours, generated from the seed
}

func newSeedTheme(seed color.NRGBA) *seedTheme {
	palettes := theme.CorePalettesFromPrimary(seed)
	return &seedTheme{
		Theme: theme.DefaultTheme(),
		light: theme.LightSchemeFromPalettes(palettes),
		dark:  theme.DarkSchemeFromPalettes(palettes),
	}
}

func (t *seedTheme) Color(n bellui.ThemeColorName, v bellui.ThemeVariant) color.Color {
	scheme := t.light
	if v == theme.VariantDark {
		scheme = t.dark
	}
	if c := theme.ColorFromMD3Scheme(scheme, n, v); c != nil {
		return c
	}
	return t.Theme.Color(n, v) // names the scheme does not cover, such as warning
}

func (t *seedTheme) RenderStyle() theme.RenderStyle { return theme.RenderStyleModern }

//docs:end seed

func main() {
	a := app.New()
	//docs:start apply
	a.Settings().SetTheme(newSeedTheme(color.NRGBA{R: 0x00, G: 0x7a, B: 0x6e, A: 0xff}))
	//docs:end apply
	w := a.NewWindow("Seed colour")

	ok := widget.NewButton("Continue", func() {})
	ok.Importance = widget.HighImportance
	w.SetContent(container.NewVBox(
		widget.NewCard("Teal", "Generated from #007a6e", widget.NewEntry()),
		widget.NewCheck("Remember me", nil),
		ok,
	))
	bellui.ShowAndRun(w)
}
