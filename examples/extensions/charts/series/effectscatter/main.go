// An effect scatter: points with a ripple ring, over an ordinary scatter.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Effect scatter series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start effectscatter
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Sensor readings", SubText: "Alerts highlighted"},
		XAxis: []axis.Config{{Type: axis.TypeValue}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.ScatterSeries{
				SeriesName: "Readings",
				Data: []float64{
					1, 20, 2, 22, 3, 21, 4, 24, 5, 23, 6, 25, 7, 24, 8, 26,
					9, 27, 10, 26, 11, 28, 12, 27, 13, 29, 14, 30, 15, 29,
				},
				Symbol: series.SymbolCircle,
			},
			&series.EffectScatterSeries{
				SeriesName:  "Alerts",
				Data:        []float64{4, 34, 9, 37, 13, 38}, // x0, y0, x1, y1, ...
				SymbolSize:  12,
				RippleScale: 3,
			},
		},
	})
	//docs:end effectscatter

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
