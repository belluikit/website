// Lines between places, with an arrow at the destination, over a tile map.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Geo lines series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start geolines
	center := [2]float64{4.5, 48.5}
	zoom := float32(4)
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Flights from Paris"},
		Series: []series.Series{
			&series.MapSeries{SeriesName: "Map", Center: center, Zoom: zoom},
			&series.GeoLinesSeries{
				SeriesName: "Flights",
				MapCenter:  center,
				MapZoom:    zoom,
				Curved:     true,
				Width:      2.5,
				Color:      color.NRGBA{R: 220, G: 50, B: 70, A: 230},
				Data: []series.GeoLine{
					{From: [2]float64{2.35, 48.86}, To: [2]float64{-0.13, 51.51}}, // London
					{From: [2]float64{2.35, 48.86}, To: [2]float64{13.40, 52.52}}, // Berlin
					{From: [2]float64{2.35, 48.86}, To: [2]float64{-3.70, 40.42}}, // Madrid
					{From: [2]float64{2.35, 48.86}, To: [2]float64{12.50, 41.90}}, // Rome
				},
			},
		},
	})
	//docs:end geolines

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
