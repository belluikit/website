// A gauge showing one value against coloured ranges.
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
	w := a.NewWindow("Gauge series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start gauge
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Disk I/O"},
		Series: []series.Series{
			&series.GaugeSeries{
				SeriesName: "MB/s",
				Value:      312,
				Min:        0,
				Max:        500,
				ShowValue:  true,
				// Minimum at the bottom right, maximum at the bottom left,
				// with the gap at the bottom. Values grow counter-clockwise.
				StartAngle: 135,
				EndAngle:   -135,
				Sectors: []series.GaugeSector{
					{Min: 0, Max: 300, Color: color.NRGBA{R: 38, G: 186, B: 121, A: 160}},
					{Min: 300, Max: 420, Color: color.NRGBA{R: 250, G: 200, B: 88, A: 160}},
					{Min: 420, Max: 500, Color: color.NRGBA{R: 239, G: 83, B: 80, A: 160}},
				},
			},
		},
	})
	//docs:end gauge

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
