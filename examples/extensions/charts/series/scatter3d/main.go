// 3D scatter: points in space, drawn as small octahedra.
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
	w := a.NewWindow("Scatter3D series")
	w.Resize(geo.NewSize(640, 420))

	// A spiral of points, stored flat as x0, y0, z0, x1, y1, z1, ...
	var pts []float64
	for i := range 60 {
		t := float64(i) * 0.25
		pts = append(pts, 1.5*math.Cos(t), float64(i)/40, 1.5*math.Sin(t))
	}

	//docs:start scatter3d
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Spiral", SubText: "60 points"},
		Series: []series.Series{
			&series.Scatter3DSeries{
				SeriesName: "Spiral",
				Data:       pts, // x0, y0, z0, x1, y1, z1, ...
				PointSize:  0.1,
			},
		},
	})
	//docs:end scatter3d

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
