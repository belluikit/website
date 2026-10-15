// Line series: a plain line with symbols, a dashed line, and a stepped line
// with a gap where a value is missing.
package main

import (
	"math"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Line series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start line
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug"}
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Visitors", SubText: "Thousands per month"},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: months}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.LineSeries{
				SeriesName: "Web",
				Data:       []float64{82, 93, 90, 104, 129, 133, 131, 142},
				Symbol:     series.SymbolCircle,
			},
			&series.LineSeries{
				SeriesName: "Target",
				Data:       []float64{90, 95, 100, 105, 110, 115, 120, 125},
				LineStyle:  series.LineStyle{Width: 1.5, Dash: []float32{6, 4}},
			},
			&series.LineSeries{
				SeriesName: "App",
				// math.NaN() is a missing value: the line breaks there.
				Data:   []float64{40, 46, 51, math.NaN(), 63, 70, 68, 77},
				Step:   series.StepEnd,
				Symbol: series.SymbolRect,
			},
		},
	})
	//docs:end line

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
