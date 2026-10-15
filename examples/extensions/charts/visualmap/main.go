// Colouring by value with VisualMapConfig.MapColor, applied by hand.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Visual map")
	w.Resize(geo.NewSize(900, 560))

	//docs:start visualmap
	vm := component.VisualMapConfig{
		Min: 0, Max: 40,
		InRange: component.VisualMapRange{Colors: []color.Color{
			color.NRGBA{R: 255, G: 224, B: 178, A: 255}, // low
			color.NRGBA{R: 230, G: 81, B: 0, A: 255},    // high
		}},
	}

	errorsByService := map[string]float64{"auth": 4, "search": 12, "billing": 38, "media": 21, "mail": 8}
	var slices []series.PieDataItem
	for _, name := range []string{"auth", "search", "billing", "media", "mail"} {
		v := errorsByService[name]
		slices = append(slices, series.PieDataItem{Name: name, Value: v, Color: vm.MapColor(v)})
	}
	//docs:end visualmap

	c := chart.New(chart.Option{
		Title:  &chart.Title{Text: "Errors by service", SubText: "Darker slices have more errors"},
		Series: []series.Series{&series.PieSeries{SeriesName: "Errors", Data: slices, ShowLabels: true, PadAngle: 1}},
	})
	w.SetContent(c)
	bellui.ShowAndRun(w)
}
