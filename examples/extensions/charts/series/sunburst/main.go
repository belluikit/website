// A sunburst: a hierarchy drawn as rings, the root in the middle.
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
	w := a.NewWindow("Sunburst series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start sunburst
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Disk usage", SubText: "Gigabytes"},
		Series: []series.Series{
			&series.SunburstSeries{
				SeriesName: "Disk",
				Data: []series.SunburstNode{
					{Name: "Media", Children: []series.SunburstNode{
						{Name: "Video", Value: 210}, {Name: "Photos", Value: 95}, {Name: "Music", Value: 40},
					}},
					{Name: "Projects", Children: []series.SunburstNode{
						{Name: "website", Value: 12}, {Name: "bellui", Value: 30}, {Name: "datasets", Value: 88},
					}},
					{Name: "System", Value: 64},
					{Name: "Apps", Value: 51},
				},
			},
		},
	})
	//docs:end sunburst

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
