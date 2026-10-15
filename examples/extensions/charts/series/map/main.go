// A tile map. Drag to pan, scroll to zoom. Tiles come from OpenStreetMap.
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
	w := a.NewWindow("Map series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start map
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Western Europe"},
		Series: []series.Series{
			&series.MapSeries{
				SeriesName: "Map",
				Center:     [2]float64{4.5, 48.5}, // longitude, latitude
				Zoom:       4,
				// TileURL: "https://tiles.example.com/{z}/{x}/{y}.png",
			},
		},
	})
	//docs:end map

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
