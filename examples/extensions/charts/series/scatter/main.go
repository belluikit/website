// Scatter series: two groups of x, y points with different symbols.
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

// cluster returns n points around (cx, cy), stored flat as x0, y0, x1, y1, ...
func cluster(n int, cx, cy, spread float64) []float64 {
	pts := make([]float64, 0, 2*n)
	for i := range n {
		angle := float64(i) * 2.4 // spreads points around the centre
		r := spread * math.Sqrt(float64(i+1)/float64(n))
		pts = append(pts, cx+r*math.Cos(angle), cy+r*math.Sin(angle))
	}
	return pts
}

func main() {
	a := app.New()
	w := a.NewWindow("Scatter series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start scatter
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Response time vs payload", SubText: "Two server builds"},
		XAxis: []axis.Config{{Type: axis.TypeValue}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.ScatterSeries{
				SeriesName: "Build A",
				Data:       cluster(40, 30, 60, 18), // x0, y0, x1, y1, ...
				Symbol:     series.SymbolCircle,
				SymbolSize: 9,
			},
			&series.ScatterSeries{
				SeriesName: "Build B",
				Data:       cluster(40, 65, 35, 15),
				Symbol:     series.SymbolRect,
				SymbolSize: 7,
			},
		},
	})
	//docs:end scatter

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
