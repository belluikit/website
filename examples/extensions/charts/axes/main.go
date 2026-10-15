// Two Y axes, a fixed axis range, and overlays drawn in data coordinates.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Axes and overlays")
	w.Resize(geo.NewSize(800, 500))

	//docs:start overlays
	overlays := []chart.Overlay{
		&chart.ReferenceLine{Axis: chart.AxisHorizontal, Value: 200, Dashed: true, Label: "Target"},
		&chart.ReferenceBand{Axis: chart.AxisHorizontal, From: 100, To: 150, Label: "Low"},
		&chart.Annotation{X: 5, Y: 260, Text: "Best month", Arrow: true},
	}
	//docs:end overlays

	//docs:start axes
	zero, hundred := 0.0, 100.0
	c := chart.New(chart.Option{
		Title:  &chart.Title{Text: "Revenue and growth"},
		Legend: &component.LegendConfig{Position: component.LegendBottom},
		Grid:   &chart.Grid{Left: 60, Right: 50, Top: 50, Bottom: 50},
		XAxis: []axis.Config{{
			Type: axis.TypeCategory,
			Data: []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		}},
		YAxis: []axis.Config{
			{Type: axis.TypeValue, SplitCount: 5},
			{Type: axis.TypeValue, Min: &zero, Max: &hundred},
		},
		Series: []series.Series{
			&series.BarSeries{SeriesName: "Revenue", Data: []float64{150, 230, 224, 218, 135, 260}},
			&series.LineSeries{
				SeriesName: "Growth %",
				Data:       []float64{12, 53, 40, 33, 8, 77},
				YAxisIndex: 1, // plotted against the second Y axis
				Smooth:     true,
			},
		},
		Overlays: overlays,
	})
	//docs:end axes

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
