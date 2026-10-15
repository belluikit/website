// A pie chart with gaps between slices and labels outside.
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
	w := a.NewWindow("Pie series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start pie
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Storage by file type"},
		Series: []series.Series{
			&series.PieSeries{
				SeriesName: "Storage",
				Radius:     0.8,
				PadAngle:   2, // degrees of space between slices
				ShowLabels: true,
				Data: []series.PieDataItem{
					{Name: "Video", Value: 412},
					{Name: "Photos", Value: 236},
					{Name: "Documents", Value: 98},
					{Name: "Music", Value: 74},
					{Name: "Other", Value: 41},
				},
			},
		},
	})
	//docs:end pie

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
