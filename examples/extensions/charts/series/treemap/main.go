// A treemap: nested rectangles sized by value.
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
	w := a.NewWindow("Treemap series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start treemap
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Disk usage", SubText: "Gigabytes"},
		Series: []series.Series{
			&series.TreemapSeries{
				SeriesName: "Disk",
				Data: []series.TreemapNode{
					{Name: "Media", Children: []series.TreemapNode{
						{Name: "Video", Value: 210}, {Name: "Photos", Value: 95}, {Name: "Music", Value: 40},
					}},
					{Name: "Projects", Children: []series.TreemapNode{
						{Name: "website", Value: 12}, {Name: "bellui", Value: 30}, {Name: "datasets", Value: 88},
					}},
					{Name: "System", Value: 64},
					{Name: "Apps", Value: 51},
				},
			},
		},
	})
	//docs:end treemap

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
