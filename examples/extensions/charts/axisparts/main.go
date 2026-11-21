// Axis names, grid lines per axis, hidden axis parts, and two Y axes side by
// side on the right.
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
	w := a.NewWindow("Axis parts")
	w.Resize(geo.NewSize(800, 500))

	//docs:start axisparts
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Weather"},
		XAxis: []axis.Config{{
			Type: axis.TypeCategory,
			Data: []string{"Mon", "Tue", "Wed", "Thu", "Fri"},
			// On a category axis, grid lines fall between the categories.
			SplitLine: axis.SplitLineShow,
		}},
		YAxis: []axis.Config{
			{Type: axis.TypeValue, Name: "°C"}, // the first Y axis has grid lines
			{Type: axis.TypeValue, Name: "mm", Hide: axis.Parts{Ticks: true}},
			{
				Type:     axis.TypeValue,
				Name:     "km/h",
				Position: axis.PositionRight,
				Offset:   60, // beside the "mm" axis, not on top of it
				Hide:     axis.Parts{Line: true},
			},
		},
		Series: []series.Series{
			&series.LineSeries{SeriesName: "Temperature", Data: []float64{12, 15, 14, 18, 17}},
			&series.BarSeries{SeriesName: "Rain", Data: []float64{4, 0, 9, 2, 1}, YAxisIndex: 1},
			&series.LineSeries{SeriesName: "Wind", Data: []float64{20, 35, 28, 12, 18}, YAxisIndex: 2},
		},
	})
	//docs:end axisparts

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
