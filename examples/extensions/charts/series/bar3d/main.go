// 3D bars on an X-Z grid. Drag to orbit, scroll to zoom.
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
	w := a.NewWindow("Bar3D series")
	w.Resize(geo.NewSize(640, 420))

	// Sales per region (X) and year (Z).
	sales := [][]float64{
		{80, 120, 60},
		{100, 150, 90},
		{70, 110, 130},
	}

	//docs:start bar3d
	var bars []series.Bar3DData
	for z, row := range sales {
		for x, v := range row {
			bars = append(bars, series.Bar3DData{X: float64(x), Z: float64(z), Value: v})
		}
	}
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Sales by region and year"},
		Series: []series.Series{
			&series.Bar3DSeries{SeriesName: "Sales", Data: bars, BarSize: 0.5},
		},
	})
	//docs:end bar3d

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
