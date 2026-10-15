// A sankey diagram: flows from sources on the left to targets on the right.
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
	w := a.NewWindow("Sankey series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start sankey
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Where visitors go", SubText: "Sessions per week"},
		Series: []series.Series{
			&series.SankeySeries{
				SeriesName: "Traffic",
				Nodes: []series.SankeyNode{
					{Name: "Search"}, {Name: "Social"}, {Name: "Direct"}, // 0-2: sources
					{Name: "Docs"}, {Name: "Download"}, {Name: "Blog"}, // 3-5: targets
				},
				Links: []series.SankeyLink{
					{Source: 0, Target: 3, Value: 420},
					{Source: 0, Target: 4, Value: 180},
					{Source: 1, Target: 5, Value: 210},
					{Source: 1, Target: 3, Value: 60},
					{Source: 2, Target: 4, Value: 240},
					{Source: 2, Target: 3, Value: 90},
				},
			},
		},
	})
	//docs:end sankey

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
