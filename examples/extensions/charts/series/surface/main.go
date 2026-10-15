// A 3D surface: a grid of heights, coloured by height.
package main

import (
	"math"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Surface series")
	w.Resize(geo.NewSize(640, 420))

	const rows, cols = 25, 25
	heights := make([]float64, rows*cols)
	for r := range rows {
		for c := range cols {
			x := float64(c)/(cols-1)*6 - 3
			z := float64(r)/(rows-1)*6 - 3
			d := math.Hypot(x, z)
			heights[r*cols+c] = math.Cos(d*1.5) * math.Exp(-d/3)
		}
	}

	//docs:start surface
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Ripple", SubText: "cos(1.5r)·e^(−r/3)"},
		Series: []series.Series{
			&series.SurfaceSeries{
				SeriesName: "Ripple",
				Rows:       rows,
				Cols:       cols,
				Data:       heights, // row-major: Data[row*Cols+col] is the height
			},
		},
	})
	//docs:end surface

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
