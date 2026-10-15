// Parallel coordinates: each line is one item, each vertical axis one measure.
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
	w := a.NewWindow("Parallel series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start parallel
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Cars", SubText: "One line per model"},
		Series: []series.Series{
			&series.ParallelSeries{
				SeriesName: "Cars",
				Axes: []series.ParallelAxis{
					{Name: "Power (kW)", Min: 50, Max: 250},
					{Name: "Weight (kg)", Min: 900, Max: 2200},
					{Name: "0-100 (s)", Min: 4, Max: 14},
					{Name: "l/100km", Min: 3, Max: 12},
					{Name: "Price (k)", Min: 10, Max: 80},
				},
				// Data[line][axis]
				Data: [][]float64{
					{70, 1050, 12.5, 4.8, 18},
					{110, 1350, 9.1, 5.9, 27},
					{150, 1600, 7.4, 7.2, 41},
					{200, 1750, 5.6, 8.9, 62},
					{240, 1900, 4.4, 10.5, 78},
					{90, 1200, 10.8, 5.2, 22},
				},
				LineWidth: 2,
			},
		},
	})
	//docs:end parallel

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
