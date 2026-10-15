// Points at longitude and latitude, drawn over a tile map.
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
	w := a.NewWindow("Geo scatter series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start geoscatter
	center := [2]float64{4.5, 48.5}
	zoom := float32(4)
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Offices"},
		Series: []series.Series{
			&series.MapSeries{SeriesName: "Map", Center: center, Zoom: zoom},
			&series.GeoScatterSeries{
				SeriesName: "Offices",
				MapCenter:  center, // match the MapSeries
				MapZoom:    zoom,
				Data: []float64{ // lon0, lat0, lon1, lat1, ...
					-0.13, 51.51, // London
					2.35, 48.86, // Paris
					4.90, 52.37, // Amsterdam
					13.40, 52.52, // Berlin
					-3.70, 40.42, // Madrid
					12.50, 41.90, // Rome
				},
				SymbolSize: 12,
				Color:      color.NRGBA{R: 220, G: 50, B: 70, A: 255},
			},
		},
	})
	//docs:end geoscatter

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
