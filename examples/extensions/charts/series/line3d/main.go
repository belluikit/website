// 3D lines: three series placed one behind the other.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Line3D series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start line3d
	depth := func(z float64) *float64 { return &z }
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Revenue, profit and cost", SubText: "One line per depth"},
		Series: []series.Series{
			&series.Line3DSeries{
				SeriesName: "Revenue",
				// With ZPosition set, Data is x, y pairs.
				Data:      []float64{0, 0.5, 1, 2.3, 2, 1.8, 3, 3.1, 4, 2.5, 5, 4.0, 6, 3.5},
				ZPosition: depth(0),
				LineWidth: 0.2,
			},
			&series.Line3DSeries{
				SeriesName: "Profit",
				Data:       []float64{0, 0.3, 1, 1.2, 2, 0.8, 3, 2.0, 4, 1.5, 5, 2.8, 6, 2.2},
				ZPosition:  depth(1.5),
				LineWidth:  0.2,
			},
			&series.Line3DSeries{
				SeriesName: "Cost",
				Data:       []float64{0, 0.2, 1, 1.1, 2, 1.0, 3, 1.1, 4, 1.0, 5, 1.2, 6, 1.3},
				ZPosition:  depth(3),
				LineWidth:  0.2,
			},
		},
	})
	//docs:end line3d

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
