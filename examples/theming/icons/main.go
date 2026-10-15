// Theme icons: looking them up, using them in widgets, and colouring your
// own SVG icon so it follows the theme.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

//docs:start custom
var (
	// A 24x24 SVG drawn in black. A themed resource replaces its fill
	// colour with a theme colour, so it matches light and dark variants.
	starSVG = bellui.NewStaticResource("star.svg", []byte(
		`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">`+
			`<path fill="#000000" d="M12 17.3 18.2 21l-1.6-7L22 9.2l-7.2-.6L12 2 9.2 8.6 2 9.2 7.4 14l-1.6 7z"/>`+
			`</svg>`))

	star        = theme.NewThemedResource(starSVG)        // foreground colour
	primaryStar = theme.NewPrimaryThemedResource(starSVG) // primary colour
	errorStar   = theme.NewErrorThemedResource(starSVG)   // error colour
)

//docs:end custom

func main() {
	a := app.New()
	w := a.NewWindow("Icons")

	//docs:start lookup
	// Two ways to get the same icon from the current theme.
	home := theme.HomeIcon()
	settings := theme.Icon(theme.IconNameSettings)

	toolbar := container.NewHBox(
		widget.NewButtonWithIcon("Home", home, func() {}),
		widget.NewButtonWithIcon("Settings", settings, func() {}),
		widget.NewIcon(theme.Icon(theme.IconNameNavigateBack)),
	)
	//docs:end lookup

	w.SetContent(container.NewVBox(
		toolbar,
		container.NewHBox(widget.NewIcon(star), widget.NewIcon(primaryStar), widget.NewIcon(errorStar)),
	))
	bellui.ShowAndRun(w)
}
