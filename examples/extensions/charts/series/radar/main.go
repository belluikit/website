// A radar chart comparing two products on five measures.
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
	w := a.NewWindow("Radar series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start radar
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Laptop comparison", SubText: "Scores out of 10"},
		Series: []series.Series{
			&series.RadarSeries{
				SeriesName: "Laptops",
				Indicators: []series.RadarIndicator{
					{Name: "Battery", Max: 10}, {Name: "Screen", Max: 10},
					{Name: "Weight", Max: 10}, {Name: "Speed", Max: 10},
					{Name: "Price", Max: 10},
				},
				// One row per item, one value per indicator.
				Data: [][]float64{
					{8, 6, 9, 5, 7},
					{5, 9, 4, 9, 4},
				},
				ShowArea: true,
				Opacity:  0.25,
			},
		},
	})
	//docs:end radar

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
